package goeth_tx_helper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func GetPublicAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return common.Address{}, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func estimateGas(ethClient *ethclient.Client, from common.Address, to *common.Address, value *big.Int, data []byte) (gasLimit uint64, err error) {
	msg := ethereum.CallMsg{
		From:  from,
		To:    to,
		Value: value,
		Data:  data,
	}

	gasLimit, err = ethClient.EstimateGas(context.Background(), msg)

	if err != nil {
		return 0, WrapExternalError(err, "failed to estimate gas limit for given operation")
	}

	return gasLimit, nil
}
