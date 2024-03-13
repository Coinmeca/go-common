package abi

var FarmSigMap = map[string]string{
	"get_all": "880a3c86",
}

const Farm = `[
{
	"inputs": [],
	"name": "get_all",
	"outputs": [
		{
			"components": [
				{
					"internalType": "address",
					"name": "farm",
					"type": "address"
				},
				{
					"internalType": "string",
					"name": "name",
					"type": "string"
				},
				{
					"internalType": "address",
					"name": "stake",
					"type": "address"
				},
				{
					"internalType": "string",
					"name": "stake_symbol",
					"type": "string"
				},
				{
					"internalType": "string",
					"name": "stake_name",
					"type": "string"
				},
				{
					"internalType": "address",
					"name": "earn",
					"type": "address"
				},
				{
					"internalType": "string",
					"name": "earn_symbol",
					"type": "string"
				},
				{
					"internalType": "string",
					"name": "earn_name",
					"type": "string"
				},
				{
					"internalType": "uint256",
					"name": "start",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "period",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "goal",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "rewards",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "locked",
					"type": "uint256"
				}
			],
			"internalType": "struct Type.Farm[]",
			"name": "",
			"type": "tuple[]"
		}
	],
	"stateMutability": "view",
	"type": "function"
	}
]`
