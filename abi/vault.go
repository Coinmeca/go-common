package abi

var VaultSigMap = map[string]string{
	"get_all":        "880a3c86",
	"get_tokens":     "bab2b9e7",
	"get_key_tokens": "566fc4e4",
}

/*
"880a3c86": "get_all()",
"566fc4e4": "get_key_tokens()",
"fdc7a143": "get_liquidity(address,address)",
"bab2b9e7": "get_tokens()",
*/

const (
	DepositEventHash  = "dcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7"
	WithdrawEventHash = "f341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567"
	ListingEventHash  = "6935191bc626032d8c74ced321b2d6a62df2578bf60599625cb4f59d51e3849f"
)

const Vault = `[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "_address",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "_token",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "_amount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_meca",
				"type": "uint256"
			}
		],
		"name": "Deposit",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "fee",
				"type": "uint256"
			}
		],
		"name": "Fee",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "_token",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "_quantity",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "_proof",
				"type": "uint256"
			}
		],
		"name": "Listing",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "reward",
				"type": "uint256"
			}
		],
		"name": "Reward",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "_address",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "_token",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "_amount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_meca",
				"type": "uint256"
			}
		],
		"name": "Withdraw",
		"type": "event"
	},
	{
		"inputs": [],
		"name": "get_all",
		"outputs": [
			{
				"components": [
					{
						"internalType": "bool",
						"name": "key",
						"type": "bool"
					},
					{
						"internalType": "address",
						"name": "addr",
						"type": "address"
					},
					{
						"internalType": "string",
						"name": "name",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "symbol",
						"type": "string"
					},
					{
						"internalType": "uint8",
						"name": "decimals",
						"type": "uint8"
					},
					{
						"internalType": "uint256",
						"name": "treasury",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "rate",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "weight",
						"type": "uint256"
					},
					{
						"internalType": "int256",
						"name": "need",
						"type": "int256"
					}
				],
				"internalType": "struct Type.Token[]",
				"name": "",
				"type": "tuple[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "get_key_tokens",
		"outputs": [
			{
				"components": [
					{
						"internalType": "bool",
						"name": "key",
						"type": "bool"
					},
					{
						"internalType": "address",
						"name": "addr",
						"type": "address"
					},
					{
						"internalType": "string",
						"name": "name",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "symbol",
						"type": "string"
					},
					{
						"internalType": "uint8",
						"name": "decimals",
						"type": "uint8"
					},
					{
						"internalType": "uint256",
						"name": "treasury",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "rate",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "weight",
						"type": "uint256"
					},
					{
						"internalType": "int256",
						"name": "need",
						"type": "int256"
					}
				],
				"internalType": "struct Type.Token[]",
				"name": "",
				"type": "tuple[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_base",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_quote",
				"type": "address"
			}
		],
		"name": "get_liquidity",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "l",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "get_tokens",
		"outputs": [
			{
				"components": [
					{
						"internalType": "bool",
						"name": "key",
						"type": "bool"
					},
					{
						"internalType": "address",
						"name": "addr",
						"type": "address"
					},
					{
						"internalType": "string",
						"name": "name",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "symbol",
						"type": "string"
					},
					{
						"internalType": "uint8",
						"name": "decimals",
						"type": "uint8"
					},
					{
						"internalType": "uint256",
						"name": "treasury",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "rate",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "weight",
						"type": "uint256"
					},
					{
						"internalType": "int256",
						"name": "need",
						"type": "int256"
					}
				],
				"internalType": "struct Type.Token[]",
				"name": "",
				"type": "tuple[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "quantity",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "proof",
				"type": "uint256"
			}
		],
		"name": "_Deposit",
		"type": "event"
	}
]`
