package goeth_tx_helper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/anxp/array-basics"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

/*
Transaction Structure Under EIP-1559

A transaction under EIP-1559 includes the following fields:

    maxFeePerGas: The maximum total fee the user is willing to pay per unit of gas (base fee + priority fee).
    maxPriorityFeePerGas: The maximum tip the user is willing to pay to miners.
    base fee: The mandatory fee that is burned and dynamically adjusts based on network demand.
    effective gas price: The actual gas price paid by the transaction, calculated as the base fee plus the priority fee.
*/

/*
How to calculate Gas:
	https://www.blocknative.com/blog/eip-1559-fees
*/

type EIP1559TransactionHelper struct {
	rpcUrl    string
	ethClient *ethclient.Client
	gasTipCap *big.Int

	emulation   bool          // Emulation of sending instead of real sending
	receiptMock types.Receipt // If emulation enabled, SendTransaction will always return this receipt mock
}

type Gas1559Params struct {
	GasTipCap *big.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap *big.Int // a.k.a. maxFeePerGas
	Gas       uint64
}

// CreateEIP1559TxHelper
//
//	gasTip is optional parameter, if 0 passed, default gasTip = 2_000_000_000 is applied.
//	However, 2_000_000_000 is too much for L2 networks, it's much bigger than L2 base fee,
//	and when calculating total transaction cost = gasAmountNeeded * maxFeePerGas,
//	where maxFeePerGas = 2 * baseFee + GasTipCap (https://www.blocknative.com/blog/eip-1559-fees),
//	total gas needed (gasAmountNeeded) can be too high and unacceptable for low banks (greater than transaction net value),
//	and lead to error "insufficient funds for gas * price + value"
//	For L2 networks gasTip = 100_000_000 is recommended
func CreateEIP1559TxHelper(rpcUrl string, gasTip int64, emulation bool, receiptMock types.Receipt) *EIP1559TransactionHelper {

	if txHelper := createTxHelperRegistry().getTxHelper(rpcUrl); txHelper != nil {
		return txHelper
	}

	ethClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	if gasTip <= 0 {
		gasTip = 2_000_000_000
	}

	txHelper := &EIP1559TransactionHelper{
		rpcUrl:      rpcUrl,
		ethClient:   ethClient,
		gasTipCap:   big.NewInt(gasTip),
		emulation:   emulation,
		receiptMock: receiptMock,
	}

	createTxHelperRegistry().addTxHelperToRegistry(txHelper)

	return txHelper
}

func (eipHelper *EIP1559TransactionHelper) GetGasParameters(from common.Address, to *common.Address, value *big.Int, data []byte) (Gas1559Params, error) {
	if eipHelper.emulation {
		return Gas1559Params{
			GasTipCap: big.NewInt(0),
			GasFeeCap: big.NewInt(0),
			Gas:       0,
		}, nil
	}

	header, err := eipHelper.ethClient.HeaderByNumber(context.Background(), nil)

	if err != nil {
		return Gas1559Params{}, WrapExternalError(err, "failed to request last block header")
	}

	baseFee := header.BaseFee

	// Doubling the Base Fee when calculating the Max Fee ensures that your transaction will remain marketable for six consecutive 100% full blocks.
	gasFeeCap := big.NewInt(0)                                                // a.k.a. maxFeePerGas
	gasFeeCap.Mul(baseFee, big.NewInt(2)).Add(gasFeeCap, eipHelper.gasTipCap) // Calculate the max fee per gas (2*baseFee + gasTipCap)

	gasLimit, err := estimateGas(eipHelper.ethClient, from, to, value, data)

	if err != nil {
		return Gas1559Params{}, err
	}

	return Gas1559Params{
		GasTipCap: eipHelper.gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
	}, nil
}

func (eipHelper *EIP1559TransactionHelper) GetBaseFee() (*big.Int, error) {
	header, err := eipHelper.ethClient.HeaderByNumber(context.Background(), nil)

	if err != nil {
		return nil, WrapExternalError(err, "failed to request last block header")
	}

	baseFee := header.BaseFee

	return baseFee, nil
}

func (eipHelper *EIP1559TransactionHelper) SendTransaction(
	privateKey *ecdsa.PrivateKey,
	to *common.Address,
	chainID *big.Int,
	gasParams Gas1559Params,
	value *big.Int,
	data []byte,
) (receipt *types.Receipt, err error) {

	if eipHelper.emulation {
		return &eipHelper.receiptMock, nil
	}

	from, err := GetPublicAddressFromPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	nonce, err := eipHelper.ethClient.PendingNonceAt(context.Background(), from)
	if err != nil {
		return nil, WrapExternalError(err, "failed to get nonce")
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: gasParams.GasTipCap,
		GasFeeCap: gasParams.GasFeeCap,
		Gas:       gasParams.Gas,
		To:        to,
		Value:     value,
		Data:      data,
	})

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %s", err) // sign is not an external call, isn't it? so we don't use wrapper for external error
	}

	if err = eipHelper.ethClient.SendTransaction(context.Background(), signedTx); err != nil {
		// Possible errors:
		// 1. insufficient funds for gas * price + value (https://ethereum.stackexchange.com/questions/78072/get-an-error-insufficient-funds-for-gas-price-value)
		// 2. replacement transaction underpriced
		return nil, WrapExternalError(err, "failed to send transaction")
	}

	receipt, err = bind.WaitMined(context.Background(), eipHelper.ethClient, signedTx)
	if err != nil {
		return nil, WrapExternalError(err, "transaction probably has not been mined (timeout?)")
	}

	return receipt, nil
}

// FilterTransactionLog filters INDEXED (only topics) transaction logs by applying ethereum.FilterQuery filter
//
// How filter works:
//
//	BlockHash |
//	FromBlock |
//	ToBlock   |- are all ignored, because we expect either txReceipt OR txLogs from one specific transaction, we should not (and don't want to) iterate through blocks!
//
// Addresses - restricts matches to events created by specific contracts, if EMPTY -> ANY address will pass, also see -> ethereum.FilterQuery
//
// Topics - restricts matches to particular event topics, an empty element slice matches any topic, also see -> ethereum.FilterQuery
func (eipHelper *EIP1559TransactionHelper) FilterTransactionLog(txReceipt *types.Receipt, txLogs []*types.Log, filter ethereum.FilterQuery) ([]types.Log, error) {
	if txReceipt != nil && txLogs != nil {
		panic("receipt OR logs should be provided, but not both!")
	}

	if txReceipt == nil && txLogs == nil {
		return nil, fmt.Errorf("no input provided")
	}

	var logsIn []*types.Log
	var addressWhitelistHex []string
	var topicsFilterPatternHex [][]string
	logsOut := make([]types.Log, 0)

	// ======================= Decide which data we'll process =========================================================
	if txReceipt != nil {
		logsIn = txReceipt.Logs
	} else {
		logsIn = txLogs
	}
	// =================================================================================================================

	// ======================= Make addresses list more convenient (to compare) ========================================
	if len(filter.Addresses) > 0 {
		addressWhitelistHex = make([]string, len(filter.Addresses), len(filter.Addresses))

		for i, address := range filter.Addresses {
			addressWhitelistHex[i] = address.Hex()
		}
	}
	// =================================================================================================================

	// ======================= Make topic pattern more convenient (to compare) =========================================
	if len(filter.Topics) > 0 {
		topicsFilterPatternHex = make([][]string, len(filter.Topics))

		for i, topicVariations := range filter.Topics {
			topicsFilterPatternHex[i] = make([]string, len(topicVariations))

			for j, variation := range topicVariations {
				topicsFilterPatternHex[i][j] = variation.Hex()
			}
		}
	}
	// =================================================================================================================

	// ======================= Finally - check every log record against filter =========================================
	for _, log := range logsIn {
		if len(addressWhitelistHex) > 0 { // We apply filter, ONLY if it is SET, all empty filters automatically satisfied
			if array_basics.InArray(addressWhitelistHex, log.Address.Hex()) == false {
				continue // If given log does not pass by creator address - SKIP it
			}
		}

		allTopicsMatched := true
		for i, topic := range log.Topics {
			if i > len(topicsFilterPatternHex)-1 { // We don't have topic filters for all next topics, so they automatically passed
				break
			}

			if len(topicsFilterPatternHex[i]) > 0 { // We apply filter, ONLY if it is SET, all empty filters automatically satisfied
				if array_basics.InArray(topicsFilterPatternHex[i], topic.Hex()) == false {
					allTopicsMatched = false
					break
				}
			}
		}

		if allTopicsMatched {
			logsOut = append(logsOut, *log)
		}
	}
	// =================================================================================================================

	return logsOut, nil
}

// ContractFunctionCall calls method "methodName" in contract "contractAddress"
//
// [READONLY] This is NOT-state-changing call
func (eipHelper *EIP1559TransactionHelper) ContractFunctionCall(contractAddress *common.Address, contractABI abi.ABI, blockNumber *big.Int, methodName string, args ...interface{}) ([]interface{}, error) {
	request, err := contractABI.Pack(methodName, args...)
	if err != nil {
		return nil, WrapExternalError(nil, err.Error()) // original error == nil  because we did not external request, all errors are local!
	}

	msg := ethereum.CallMsg{
		To:   contractAddress,
		Data: request,
	}

	result, err := eipHelper.ethClient.CallContract(context.Background(), msg, blockNumber)
	if err != nil {
		return nil, WrapExternalError(err, fmt.Sprintf("failed to call function \"%s\" at contract \"%s\"", methodName, contractAddress))
	}

	parsedResponse, err := contractABI.Unpack(methodName, result)
	if err != nil {
		// Unpack is not an external call, so any error interprets as local!
		return nil, WrapExternalError(nil, fmt.Sprintf("failed to parse response for contract function \"%s\", check contract ABI; error: %s", methodName, err))
	}

	return parsedResponse, nil
}

// ContractFunctionCallNoArguments calls method "methodName" in contract "contractAddress", suitable for methods with NO arguments
//
// [READONLY] This is NOT-state-changing call
func (eipHelper *EIP1559TransactionHelper) ContractFunctionCallNoArguments(contractAddress *common.Address, contractABI abi.ABI, blockNumber *big.Int, methodName string) ([]interface{}, error) {
	return eipHelper.ContractFunctionCall(contractAddress, contractABI, blockNumber, methodName)
}

func (eipHelper *EIP1559TransactionHelper) GetLatestBlockNumber() (*big.Int, error) {
	blockNumber, err := eipHelper.ethClient.BlockNumber(context.Background())
	if err != nil {
		return nil, WrapExternalError(err, "failed to get latest block number")
	}

	return big.NewInt(0).SetUint64(blockNumber), nil
}

func (eipHelper *EIP1559TransactionHelper) GetEthClient() *ethclient.Client {
	return eipHelper.ethClient
}

func (eipHelper *EIP1559TransactionHelper) GetRpcUrl() string {
	return eipHelper.rpcUrl
}

func (eipHelper *EIP1559TransactionHelper) GetPublicAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	return GetPublicAddressFromPrivateKey(privateKey)
}
