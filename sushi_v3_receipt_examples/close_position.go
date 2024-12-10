package sushi_v3_receipt_examples

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

// CloseETHUSDCReceiptExample - Typical receipt you receive when closing ETH+USDC pool position.
// Recreated from real transaction: https://basescan.org/tx/0x5ac58019d20d412cb9cc6ed6d7218d0416b45f801ef2a1bfc1614c8b4b0ffd47#eventlog
var CloseETHUSDCReceiptExample = types.Receipt{
	Type:              2,
	PostState:         nil,
	Status:            1,
	CumulativeGasUsed: 365592,
	Bloom:             types.Bloom{},
	Logs: []*types.Log{
		{
			Address: common.HexToAddress("0x57713F7716e0b0F65ec116912F834E49805480d2"), // Address of contract, created the transaction (BASE ETH/USDC-005 Pool)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Burn(address,int24,int24,uint128,uint256,uint256)")),
				common.HexToHash("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de"), // owner
				common.HexToHash("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffcf36a"), // tickLower
				common.HexToHash("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd01ac"), // tickUpper
			},
			// What encoded in Data:
			// 	amount:  457623340164    (amount of liquidity to burn)
			// 	amount0: 789072362800965 (amount0 - amount of token0 owed)
			// 	amount1: 201823          (amount1 - amount of token1 owed)
			// To decode - split by 32 bytes and parse to big.Int, big.Int, big.Int
			Data:        common.FromHex("0x0000000000000000000000000000000000000000000000000000006a8c7a04840000000000000000000000000000000000000000000000000002cda839e3074500000000000000000000000000000000000000000000000000000000001ecbbd"),
			BlockNumber: 19163726,
			TxHash:      common.HexToHash("0x5ac58019d20d412cb9cc6ed6d7218d0416b45f801ef2a1bfc1614c8b4b0ffd47"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x2a07791dc22682dab6f458df8a8c7cf16fb4dc1da39629085783952094fc4b6c"),
			Index:       0,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x80C7DD17B01855a6D2347444a0FCC36136a314de"), // Address of contract, created the transaction (NFT Position Manager)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("DecreaseLiquidity(uint256,uint128,uint256,uint256)")),
				common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000e2cc"), // tokenId
			},
			// What encoded in Data:
			// 	liquidity: 457623340164
			// 	amount0:   789072362800965
			// 	amount1:   2018237
			// To decode - split by 32 bytes and parse to big.Int, big.Int, big.Int
			Data:        common.FromHex("0x0000000000000000000000000000000000000000000000000000006a8c7a04840000000000000000000000000000000000000000000000000002cda839e3074500000000000000000000000000000000000000000000000000000000001ecbbd"),
			BlockNumber: 19163726,
			TxHash:      common.HexToHash("0x5ac58019d20d412cb9cc6ed6d7218d0416b45f801ef2a1bfc1614c8b4b0ffd47"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x2a07791dc22682dab6f458df8a8c7cf16fb4dc1da39629085783952094fc4b6c"),
			Index:       1,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x4200000000000000000000000000000000000006"), // Address of contract, created the transaction (Wrapped Ether )
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")),
				common.HexToHash("0x00000000000000000000000057713f7716e0b0f65ec116912f834e49805480d2"), // src
				common.HexToHash("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de"), // dst
			},
			Data:        common.FromHex("0x0000000000000000000000000000000000000000000000000002ce4f7f977b07"), // wad: 789790791793415
			BlockNumber: 19163726,
			TxHash:      common.HexToHash("0x5ac58019d20d412cb9cc6ed6d7218d0416b45f801ef2a1bfc1614c8b4b0ffd47"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x2a07791dc22682dab6f458df8a8c7cf16fb4dc1da39629085783952094fc4b6c"),
			Index:       2,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"), // Address of contract, created the transaction (Circle: USDC Token)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")),
				common.HexToHash("0x00000000000000000000000057713f7716e0b0f65ec116912f834e49805480d2"), // from
				common.HexToHash("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de"), // to
			},
			Data:        common.FromHex("0x00000000000000000000000000000000000000000000000000000000001ed2c7"), // amount of USDC: 2.020.039
			BlockNumber: 19163726,
			TxHash:      common.HexToHash("0x5ac58019d20d412cb9cc6ed6d7218d0416b45f801ef2a1bfc1614c8b4b0ffd47"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x2a07791dc22682dab6f458df8a8c7cf16fb4dc1da39629085783952094fc4b6c"),
			Index:       3,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x57713F7716e0b0F65ec116912F834E49805480d2"), // Address of contract, created the transaction (BASE ETH/USDC-005 Pool)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Collect(address,address,int24,int24,uint128,uint128)")),
				common.HexToHash("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de"), // owner
				common.HexToHash("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffcf36a"), // tickLower
				common.HexToHash("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd01ac"), // tickUpper
			},
			// What encoded in Data:
			// 	recipient: 0x80C7DD17B01855a6D2347444a0FCC36136a314de
			// 	amount0:   789790791793415
			// 	amount1:   2020039
			// To decode - split by 32 bytes and parse to address, big.Int, big.Int
			Data:        common.FromHex("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de0000000000000000000000000000000000000000000000000002ce4f7f977b0700000000000000000000000000000000000000000000000000000000001ed2c7"),
			BlockNumber: 19163726,
			TxHash:      common.HexToHash("0x5ac58019d20d412cb9cc6ed6d7218d0416b45f801ef2a1bfc1614c8b4b0ffd47"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x2a07791dc22682dab6f458df8a8c7cf16fb4dc1da39629085783952094fc4b6c"),
			Index:       4,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x80C7DD17B01855a6D2347444a0FCC36136a314de"), // Address of contract, created the transaction (NFT Position Manager)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Collect(uint256,address,uint256,uint256)")),
				common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000e2cc"), // tokenId
			},
			// What encoded in Data:
			// 	recipient: 0x80C7DD17B01855a6D2347444a0FCC36136a314de
			// 	amount0:   789790791793415
			// 	amount1:   2020039
			// To decode - split by 32 bytes and parse to address, big.Int, big.Int
			Data:        common.FromHex("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de0000000000000000000000000000000000000000000000000002ce4f7f977b0700000000000000000000000000000000000000000000000000000000001ed2c7"),
			BlockNumber: 19163726,
			TxHash:      common.HexToHash("0x5ac58019d20d412cb9cc6ed6d7218d0416b45f801ef2a1bfc1614c8b4b0ffd47"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x2a07791dc22682dab6f458df8a8c7cf16fb4dc1da39629085783952094fc4b6c"),
			Index:       5,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x4200000000000000000000000000000000000006"), // Address of contract, created the transaction (WrappedEther)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Withdrawal(address,uint256)")),
				common.HexToHash("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de"),
			},
			Data:        common.FromHex("0x0000000000000000000000000000000000000000000000000002ce4f7f977b07"),
			BlockNumber: 19163726,
			TxHash:      common.HexToHash("0x5ac58019d20d412cb9cc6ed6d7218d0416b45f801ef2a1bfc1614c8b4b0ffd47"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x2a07791dc22682dab6f458df8a8c7cf16fb4dc1da39629085783952094fc4b6c"),
			Index:       6,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"), // Address of contract, created the transaction (Circle: USDC Token)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")),
				common.HexToHash("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de"), // from
				common.HexToHash("0x00000000000000000000000035976f39bce40ce858fb66360c49231e6b8ee4a1"), // to (wallet address)
			},
			Data:        common.FromHex("0x00000000000000000000000000000000000000000000000000000000001ed2c7"), // amount of usdc to transfer, 2.020.039 ~ $2.02
			BlockNumber: 17298440,
			TxHash:      common.HexToHash("0x5ac58019d20d412cb9cc6ed6d7218d0416b45f801ef2a1bfc1614c8b4b0ffd47"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x2a07791dc22682dab6f458df8a8c7cf16fb4dc1da39629085783952094fc4b6c"),
			Index:       7,
			Removed:     false,
		},
	},
	TxHash:            common.HexToHash("0x5ac58019d20d412cb9cc6ed6d7218d0416b45f801ef2a1bfc1614c8b4b0ffd47"),
	ContractAddress:   common.Address{},
	GasUsed:           321753,
	EffectiveGasPrice: big.NewInt(2006120122),
	BlobGasUsed:       0,
	BlobGasPrice:      nil,
	BlockHash:         common.HexToHash("0x2a07791dc22682dab6f458df8a8c7cf16fb4dc1da39629085783952094fc4b6c"),
	BlockNumber:       big.NewInt(19163726),
	TransactionIndex:  1,
}

// CloseUSDCaxlUSDCReceiptExample - Typical receipt you receive when closing USDC+axlUSDC pool position.
// Recreated from real transaction: https://basescan.org/tx/0x625af890bcf8b552b59185dd97a20c559b84bb0c2c98ea6688b187b3098d0a32#eventlog
var CloseUSDCaxlUSDCReceiptExample = types.Receipt{
	Type:              2,
	PostState:         nil,
	Status:            1,
	CumulativeGasUsed: 274898,
	Bloom:             types.Bloom{},
	Logs: []*types.Log{
		{
			Address: common.HexToAddress("0x9EFe71955931C31c1c63395Df11063289332bF8b"), // Address of contract, created the transaction (BASE USDC/axlUSDC-100 Pool)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Burn(address,int24,int24,uint128,uint256,uint256)")),
				common.HexToHash("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de"), // owner
				common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000016"), // tickLower
				common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000020"), // tickUpper
			},
			// What encoded in Data:
			// 	amount:  3789253116 (amount of liquidity to burn)
			// 	amount0: 1891605    (amount0 - amount of token0 owed)
			// 	amount1: 371        (amount1 - amount of token1 owed)
			// To decode - split by 32 bytes and parse to big.Int, big.Int, big.Int
			Data:        common.FromHex("0x00000000000000000000000000000000000000000000000000000000e1db69fc00000000000000000000000000000000000000000000000000000000001cdd150000000000000000000000000000000000000000000000000000000000000173"),
			BlockNumber: 19450656,
			TxHash:      common.HexToHash("0x625af890bcf8b552b59185dd97a20c559b84bb0c2c98ea6688b187b3098d0a32"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0xe615a4aacc14fc9caec6fbfce34eb2b813daaa305f9d0bb78c936081d3c3021a"),
			Index:       0,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x80C7DD17B01855a6D2347444a0FCC36136a314de"), // Address of contract, created the transaction (NFT Position Manager)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("DecreaseLiquidity(uint256,uint128,uint256,uint256)")),
				common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000e3df"), // tokenId
			},
			// What encoded in Data:
			// 	liquidity: 3789253116
			// 	amount0:   1891605
			// 	amount1:   371
			// To decode - split by 32 bytes and parse to big.Int, big.Int, big.Int
			Data:        common.FromHex("0x00000000000000000000000000000000000000000000000000000000e1db69fc00000000000000000000000000000000000000000000000000000000001cdd150000000000000000000000000000000000000000000000000000000000000173"),
			BlockNumber: 19450656,
			TxHash:      common.HexToHash("0x625af890bcf8b552b59185dd97a20c559b84bb0c2c98ea6688b187b3098d0a32"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0xe615a4aacc14fc9caec6fbfce34eb2b813daaa305f9d0bb78c936081d3c3021a"),
			Index:       1,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"), // Address of contract, created the transaction (Circle: USDC Token)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")),
				common.HexToHash("0x0000000000000000000000009efe71955931c31c1c63395df11063289332bf8b"), // from
				common.HexToHash("0x00000000000000000000000035976f39bce40ce858fb66360c49231e6b8ee4a1"), // to
			},
			Data:        common.FromHex("0x00000000000000000000000000000000000000000000000000000000001cdd78"), // value: 1891704
			BlockNumber: 19450656,
			TxHash:      common.HexToHash("0x625af890bcf8b552b59185dd97a20c559b84bb0c2c98ea6688b187b3098d0a32"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0xe615a4aacc14fc9caec6fbfce34eb2b813daaa305f9d0bb78c936081d3c3021a"),
			Index:       2,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0xEB466342C4d449BC9f53A865D5Cb90586f405215"), // Address of contract, created the transaction (Bridged USD Coin: axlUSDC Token)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")),
				common.HexToHash("0x0000000000000000000000009efe71955931c31c1c63395df11063289332bf8b"), // from
				common.HexToHash("0x00000000000000000000000035976f39bce40ce858fb66360c49231e6b8ee4a1"), // to
			},
			Data:        common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000173"), // value: 371
			BlockNumber: 19450656,
			TxHash:      common.HexToHash("0x625af890bcf8b552b59185dd97a20c559b84bb0c2c98ea6688b187b3098d0a32"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0xe615a4aacc14fc9caec6fbfce34eb2b813daaa305f9d0bb78c936081d3c3021a"),
			Index:       3,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x9EFe71955931C31c1c63395Df11063289332bF8b"), // Address of contract, created the transaction (BASE USDC/axlUSDC-100 Pool)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Collect(address,address,int24,int24,uint128,uint128)")),
				common.HexToHash("0x00000000000000000000000080c7dd17b01855a6d2347444a0fcc36136a314de"), // owner
				common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000016"), // tickLower
				common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000020"), // tickUpper
			},
			// What encoded in Data:
			// 	recipient: 0x35976f39BCe40Ce858fB66360c49231E6B8Ee4A1
			// 	amount0:   1891704
			// 	amount1:   371
			// To decode - split by 32 bytes and parse to address, big.Int, big.Int
			Data:        common.FromHex("0x00000000000000000000000035976f39bce40ce858fb66360c49231e6b8ee4a100000000000000000000000000000000000000000000000000000000001cdd780000000000000000000000000000000000000000000000000000000000000173"),
			BlockNumber: 19450656,
			TxHash:      common.HexToHash("0x625af890bcf8b552b59185dd97a20c559b84bb0c2c98ea6688b187b3098d0a32"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0xe615a4aacc14fc9caec6fbfce34eb2b813daaa305f9d0bb78c936081d3c3021a"),
			Index:       4,
			Removed:     false,
		},
		{
			Address: common.HexToAddress("0x80C7DD17B01855a6D2347444a0FCC36136a314de"), // Address of contract, created the transaction (NFT Position Manager)
			Topics: []common.Hash{
				crypto.Keccak256Hash([]byte("Collect(uint256,address,uint256,uint256)")),
				common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000e3df"), // tokenId
			},
			// What encoded in Data:
			// 	recipient: 0x35976f39BCe40Ce858fB66360c49231E6B8Ee4A1
			// 	amount0:   1891704
			// 	amount1:   371
			// To decode - split by 32 bytes and parse to address, big.Int, big.Int
			Data:        common.FromHex("0x00000000000000000000000035976f39bce40ce858fb66360c49231e6b8ee4a100000000000000000000000000000000000000000000000000000000001cdd780000000000000000000000000000000000000000000000000000000000000173"),
			BlockNumber: 19450656,
			TxHash:      common.HexToHash("0x625af890bcf8b552b59185dd97a20c559b84bb0c2c98ea6688b187b3098d0a32"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0xe615a4aacc14fc9caec6fbfce34eb2b813daaa305f9d0bb78c936081d3c3021a"),
			Index:       5,
			Removed:     false,
		},
	},
	TxHash:            common.HexToHash("0x625af890bcf8b552b59185dd97a20c559b84bb0c2c98ea6688b187b3098d0a32"),
	ContractAddress:   common.Address{},
	GasUsed:           231059,
	EffectiveGasPrice: big.NewInt(2002626946),
	BlobGasUsed:       0,
	BlobGasPrice:      nil,
	BlockHash:         common.HexToHash("0xe615a4aacc14fc9caec6fbfce34eb2b813daaa305f9d0bb78c936081d3c3021a"),
	BlockNumber:       big.NewInt(19450656),
	TransactionIndex:  1,
}
