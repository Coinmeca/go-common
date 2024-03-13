package abi

var OrderbookSigMap = map[string]string{
	"get_asks":      "03f67026",
	"get_bids":      "4ad40b34",
	"get_orderbook": "d333c72d",
	"get_info":      "60583488",
}

/*
"60583488": "get_info()",
"03f67026": "get_asks(uint16)",
"4ad40b34": "get_bids(uint16)",
"60fa1490": "get_nft()",
"d333c72d": "get_orderbook(uint8)",
*/
const (
	BuyEventHash    = "00f93dbdb72854b6b6fb35433086556f2635fc83c37080c667496fecfa650fb4"
	SellEventHash   = "01fbb57444511e3de5b26ac09ad6bec45c3f9a1e59dd4a0f2b13a240d18476ce"
	LiquidEventHash = "b6ed19cdf32fe732da504f7525ffbec01e43cf56a407bdcda3ab1120d756a259"
	CancelEventHash = "13f70c3891bf9326a78bedbd802efd57fa4451e8b0eab22d638b4e6b7a878eaf"
	BidEventHash    = "d21fbaad97462831ad0c216f300fefb33a10b03bb18bb70ed668562e88d15d53"
	AskEventHash    = "c96f568c4dd67d35b7a2550cf8f57f2c8f156325f4531215a4bf6df0854352e3"
)

const Orderbook = `[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "_owner",
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
				"name": "_price",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_amount",
				"type": "uint256"
			}
		],
		"name": "Ask",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "_owner",
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
				"name": "_price",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_amount",
				"type": "uint256"
			}
		],
		"name": "Bid",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "_owner",
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
				"name": "_price",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_amount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_quantity",
				"type": "uint256"
			}
		],
		"name": "Buy",
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
				"name": "_price",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "_amount",
				"type": "uint256"
			}
		],
		"name": "Cancel",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "_owner",
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
				"name": "_price",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_amount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_quantity",
				"type": "uint256"
			}
		],
		"name": "Claim",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "_owner",
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
			}
		],
		"name": "Liquidation",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "_owner",
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
				"name": "_price",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_amount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_quantity",
				"type": "uint256"
			}
		],
		"name": "Sell",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "uint16",
				"name": "_range",
				"type": "uint16"
			}
		],
		"name": "get_asks",
		"outputs": [
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "price",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "balance",
						"type": "uint256"
					}
				],
				"internalType": "struct Type.Tick[]",
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
				"internalType": "uint16",
				"name": "_range",
				"type": "uint16"
			}
		],
		"name": "get_bids",
		"outputs": [
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "price",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "balance",
						"type": "uint256"
					}
				],
				"internalType": "struct Type.Tick[]",
				"name": "",
				"type": "tuple[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "get_info",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			},
			{
				"internalType": "uint8",
				"name": "",
				"type": "uint8"
			},
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "get_nft",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint8",
				"name": "_range",
				"type": "uint8"
			}
		],
		"name": "get_orderbook",
		"outputs": [
			{
				"components": [
					{
						"components": [
							{
								"internalType": "uint256",
								"name": "price",
								"type": "uint256"
							},
							{
								"internalType": "uint256",
								"name": "balance",
								"type": "uint256"
							}
						],
						"internalType": "struct Type.Tick[]",
						"name": "asks",
						"type": "tuple[]"
					},
					{
						"components": [
							{
								"internalType": "uint256",
								"name": "price",
								"type": "uint256"
							},
							{
								"internalType": "uint256",
								"name": "balance",
								"type": "uint256"
							}
						],
						"internalType": "struct Type.Tick[]",
						"name": "bids",
						"type": "tuple[]"
					}
				],
				"internalType": "struct Type.Orderbook",
				"name": "",
				"type": "tuple"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`

/*

 */
