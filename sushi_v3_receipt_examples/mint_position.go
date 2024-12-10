package sushi_v3_receipt_examples

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

// MintUSDCaxlUSDCReceiptExample - Typical receipt you receive when minting position of ERC20+ERC20 token (No native!)
// Recreated from real transaction: https://basescan.org/tx/0x95c53b8314e782b73942f9bab9afaefbcef5b44d88b3ee40c764d6b0ae41646d#eventlog
var MintUSDCaxlUSDCReceiptExample = types.Receipt{
	Type:              2,
	PostState:         nil,
	Status:            1,
	CumulativeGasUsed: 535880,
	Bloom:             types.Bloom{},
	Logs: []*types.Log{
		{
			Address: common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"), // Address of contract, created the transaction (Circle: USDC Token)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")),
				common.HexToHash("0x00000000000000000000000035976f39bce40ce858fb66360c49231e6b8ee4a1"), // from
				common.HexToHash("0x0000000000000000000000009efe71955931c31c1c63395df11063289332bf8b"), // to (BASE USDC/axlUSDC-100 Pool)
			},
			Data:        common.FromHex("0x00000000000000000000000000000000000000000000000000000000000d9b40"), // 891712 ($0.89)
			BlockNumber: 19206086,
			TxHash:      common.HexToHash("0x95c53b8314e782b73942f9bab9afaefbcef5b44d88b3ee40c764d6b0ae41646d"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x6c10cbca8a2ec7fc0a26a36c7b12d34b49d77dd85c1e3484791a4c29c3f3b532"),
			Index:       0,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0xEB466342C4d449BC9f53A865D5Cb90586f405215"), // Address of contract, created the transaction (Bridged USD Coin: axlUSDC Token)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Approval(address,address,uint256)")),
				common.HexToHash("0x00000000000000000000000035976f39bce40ce858fb66360c49231e6b8ee4a1"), // owner
				common.HexToHash("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de"), // spender
			},
			Data:        common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000000"), // 0
			BlockNumber: 19206086,
			TxHash:      common.HexToHash("0x95c53b8314e782b73942f9bab9afaefbcef5b44d88b3ee40c764d6b0ae41646d"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x6c10cbca8a2ec7fc0a26a36c7b12d34b49d77dd85c1e3484791a4c29c3f3b532"),
			Index:       1,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0xEB466342C4d449BC9f53A865D5Cb90586f405215"), // Address of contract, created the transaction (Bridged USD Coin: axlUSDC Token)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")),
				common.HexToHash("0x00000000000000000000000035976f39bce40ce858fb66360c49231e6b8ee4a1"), // from
				common.HexToHash("0x0000000000000000000000009efe71955931c31c1c63395df11063289332bf8b"), // to
			},
			Data:        common.FromHex("0x00000000000000000000000000000000000000000000000000000000000f4cec"), // 1002732 ($1)
			BlockNumber: 19206086,
			TxHash:      common.HexToHash("0x95c53b8314e782b73942f9bab9afaefbcef5b44d88b3ee40c764d6b0ae41646d"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x6c10cbca8a2ec7fc0a26a36c7b12d34b49d77dd85c1e3484791a4c29c3f3b532"),
			Index:       2,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x9EFe71955931C31c1c63395Df11063289332bF8b"), // Address of contract, created the transaction (BASE USDC/axlUSDC-100 Pool)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Mint(address,address,int24,int24,uint128,uint256,uint256)")),
				common.HexToHash("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de"), // owner
				common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000016"), // tickLower
				common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000020"), // tickUpper
			},
			// What encoded in Data:
			// sender:  0x80C7DD17B01855a6D2347444a0FCC36136a314de
			// amount:  3789253116
			// amount0: 891712
			// amount1: 1002732
			Data:        common.FromHex("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de00000000000000000000000000000000000000000000000000000000e1db69fc00000000000000000000000000000000000000000000000000000000000d9b4000000000000000000000000000000000000000000000000000000000000f4cec"),
			BlockNumber: 19206086,
			TxHash:      common.HexToHash("0x95c53b8314e782b73942f9bab9afaefbcef5b44d88b3ee40c764d6b0ae41646d"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x6c10cbca8a2ec7fc0a26a36c7b12d34b49d77dd85c1e3484791a4c29c3f3b532"),
			Index:       3,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x80C7DD17B01855a6D2347444a0FCC36136a314de"), // Address of contract, created the transaction (BASE NFT Position Manager)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")),
				common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"), // from
				common.HexToHash("0x00000000000000000000000035976f39bce40ce858fb66360c49231e6b8ee4a1"), // to
				common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000e3df"), // tokenId
			},
			Data:        []byte{},
			BlockNumber: 19206086,
			TxHash:      common.HexToHash("0x95c53b8314e782b73942f9bab9afaefbcef5b44d88b3ee40c764d6b0ae41646d"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x6c10cbca8a2ec7fc0a26a36c7b12d34b49d77dd85c1e3484791a4c29c3f3b532"),
			Index:       4,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x80C7DD17B01855a6D2347444a0FCC36136a314de"), // Address of contract, created the transaction (BASE NFT Position Manager)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("IncreaseLiquidity(uint256,uint128,uint256,uint256)")),
				common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000e3df"), // tokenId (58335)
			},
			// What in Data:
			// liquidity: 3789253116
			// amount0:   891712
			// amount1:   1002732
			Data:        common.FromHex("0x00000000000000000000000000000000000000000000000000000000e1db69fc00000000000000000000000000000000000000000000000000000000000d9b4000000000000000000000000000000000000000000000000000000000000f4cec"),
			BlockNumber: 19206086,
			TxHash:      common.HexToHash("0x95c53b8314e782b73942f9bab9afaefbcef5b44d88b3ee40c764d6b0ae41646d"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x6c10cbca8a2ec7fc0a26a36c7b12d34b49d77dd85c1e3484791a4c29c3f3b532"),
			Index:       5,
			Removed:     false,
		},
	},
	TxHash:            common.HexToHash("0x95c53b8314e782b73942f9bab9afaefbcef5b44d88b3ee40c764d6b0ae41646d"),
	ContractAddress:   common.Address{},
	GasUsed:           492053,
	EffectiveGasPrice: big.NewInt(2006033430),
	BlobGasUsed:       0,
	BlobGasPrice:      nil,
	BlockHash:         common.HexToHash("0x6c10cbca8a2ec7fc0a26a36c7b12d34b49d77dd85c1e3484791a4c29c3f3b532"),
	BlockNumber:       big.NewInt(19206086),
	TransactionIndex:  1,
}

// MintETHUSDCReceiptExample - Typical receipt you receive when minting pool position with native token (Eth) and ERC20 one.
// Recreated from real transaction: https://basescan.org/tx/0xbcb37d85c1850e923e5b4238797e84141cea9dedd327aaa74cf76cb4dc3bdbf4#eventlog
var MintETHUSDCReceiptExample = types.Receipt{
	Type:              2,
	PostState:         nil,
	Status:            1,
	CumulativeGasUsed: 488471,
	Bloom:             types.Bloom{},
	Logs: []*types.Log{
		{
			Address: common.HexToAddress("0x4200000000000000000000000000000000000006"), // Address of contract, created the transaction (Wrapped Ether)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Deposit(address,uint256)")),
				common.HexToHash("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de"), // dst
			},
			Data:        common.FromHex("0x0000000000000000000000000000000000000000000000000005765d6ffc7305"), // wad: 1537518566404869
			BlockNumber: 17298440,
			TxHash:      common.HexToHash("0xbcb37d85c1850e923e5b4238797e84141cea9dedd327aaa74cf76cb4dc3bdbf4"),
			TxIndex:     1,
			BlockHash:   common.Hash{},
			Index:       0,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x4200000000000000000000000000000000000006"), // Address of contract, created the transaction (Wrapped Ether)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")),
				common.HexToHash("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de"), // src
				common.HexToHash("0x00000000000000000000000057713f7716e0b0f65ec116912f834e49805480d2"), // dst
			},
			Data:        common.FromHex("0x0000000000000000000000000000000000000000000000000005765d6ffc7305"), // wad: 1537518566404869
			BlockNumber: 17298440,
			TxHash:      common.HexToHash("0xbcb37d85c1850e923e5b4238797e84141cea9dedd327aaa74cf76cb4dc3bdbf4"),
			TxIndex:     1,
			BlockHash:   common.Hash{},
			Index:       1,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"), // Address of contract, created the transaction (Circle: USDC Token)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")),
				common.HexToHash("0x00000000000000000000000035976f39bce40ce858fb66360c49231e6b8ee4a1"), // from
				common.HexToHash("0x00000000000000000000000057713f7716e0b0f65ec116912f834e49805480d2"), // to
			},
			Data:        common.FromHex("0x000000000000000000000000000000000000000000000000000000000051d610"), // value: 5363216
			BlockNumber: 17298440,
			TxHash:      common.HexToHash("0xbcb37d85c1850e923e5b4238797e84141cea9dedd327aaa74cf76cb4dc3bdbf4"),
			TxIndex:     1,
			BlockHash:   common.Hash{},
			Index:       2,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x57713F7716e0b0F65ec116912F834E49805480d2"), // Address of contract, created the transaction ()
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Mint(address,address,int24,int24,uint128,uint256,uint256)")),
				common.HexToHash("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de"), // owner
				common.HexToHash("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd0616"), // tickLower (-195050)
				common.HexToHash("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd067a"), // tickUpper (-194950)
			},
			// What in Data:
			// sender:  0x80C7DD17B01855a6D2347444a0FCC36136a314de
			// amount:  36373200899982
			// amount0: 1537518566404869
			// amount1: 5363216
			Data:        common.FromHex("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de00000000000000000000000000000000000000000000000000002114cbb80f8e0000000000000000000000000000000000000000000000000005765d6ffc7305000000000000000000000000000000000000000000000000000000000051d610"),
			BlockNumber: 17298440,
			TxHash:      common.HexToHash("0xbcb37d85c1850e923e5b4238797e84141cea9dedd327aaa74cf76cb4dc3bdbf4"),
			TxIndex:     1,
			BlockHash:   common.Hash{},
			Index:       3,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x80C7DD17B01855a6D2347444a0FCC36136a314de"), // Address of contract, created the transaction (NonfungiblePositionManager)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")),
				common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
				common.HexToHash("0x35976f39BCe40Ce858fB66360c49231E6B8Ee4A1"), // Recipient address
				common.BytesToHash(big.NewInt(51230).Bytes()),
			},
			Data:        []byte{},
			BlockNumber: 17298440,
			TxHash:      common.HexToHash("0xbcb37d85c1850e923e5b4238797e84141cea9dedd327aaa74cf76cb4dc3bdbf4"),
			TxIndex:     1,
			BlockHash:   common.Hash{},
			Index:       4,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x80C7DD17B01855a6D2347444a0FCC36136a314de"), // Address of contract, created the transaction (NonfungiblePositionManager)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("IncreaseLiquidity(uint256,uint128,uint256,uint256)")),
				common.BytesToHash(big.NewInt(51230).Bytes()), // tokenId
			},
			// What in Data:
			// liquidity: 36373200899982
			// amount0:   1537518566404869
			// amount1:   5363216
			Data:        common.FromHex("0x00000000000000000000000000000000000000000000000000002114cbb80f8e0000000000000000000000000000000000000000000000000005765d6ffc7305000000000000000000000000000000000000000000000000000000000051d610"),
			BlockNumber: 17298440,
			TxHash:      common.HexToHash("0xbcb37d85c1850e923e5b4238797e84141cea9dedd327aaa74cf76cb4dc3bdbf4"),
			TxIndex:     1,
			BlockHash:   common.Hash{},
			Index:       5,
			Removed:     false,
		},
	},
	TxHash:            common.HexToHash("0xbcb37d85c1850e923e5b4238797e84141cea9dedd327aaa74cf76cb4dc3bdbf4"),
	ContractAddress:   common.HexToAddress("0x80C7DD17B01855a6D2347444a0FCC36136a314de"),
	GasUsed:           444620,
	EffectiveGasPrice: big.NewInt(0),
	BlobGasUsed:       0,
	BlobGasPrice:      nil,
	BlockHash:         common.Hash{},
	BlockNumber:       big.NewInt(17298440),
	TransactionIndex:  1,
}
