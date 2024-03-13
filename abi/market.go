package abi

var MarketSigMap = map[string]string{
	"get_all":     "880a3c86",
	"get_markets": "2c163c0e",
}

/* hashes
   	"880a3c86": "get_all()",
   	"2c163c0e": "get_markets(address)"
	"f68378ad": "market(address,address)"
*/

const Market = `[
{
	"inputs": [],
    "name": "get_all",
    "outputs": [
      {
        "components": [
		{
			"internalType": "address",
			"name": "market",
			"type": "address"
		},
		{
			"internalType": "address",
			"name": "nft",
			"type": "address"
		},
          {
            "internalType": "string",
            "name": "symbol",
            "type": "string"
          },
          {
            "internalType": "string",
            "name": "name",
            "type": "string"
          },
          {
            "internalType": "address",
            "name": "base",
            "type": "address"
          },
          {
            "internalType": "address",
            "name": "quote",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "price",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "tick",
            "type": "uint256"
          },
          {
            "internalType": "uint8",
            "name": "fee",
            "type": "uint8"
          },
		{
			"internalType": "bool",
			"name": "lock",
			"type": "bool"
		}
        ],
        "internalType": "struct Type.Market[]",
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
      }
    ],
    "name": "get_markets",
    "outputs": [
      {
        "components": [
		{
			"internalType": "address",
			"name": "market",
			"type": "address"
		},
		{
			"internalType": "address",
			"name": "nft",
			"type": "address"
		},
          {
            "internalType": "string",
            "name": "symbol",
            "type": "string"
          },
          {
            "internalType": "string",
            "name": "name",
            "type": "string"
          },
          {
            "internalType": "address",
            "name": "base",
            "type": "address"
          },
          {
            "internalType": "address",
            "name": "quote",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "price",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "tick",
            "type": "uint256"
          },
          {
            "internalType": "uint8",
            "name": "fee",
            "type": "uint8"
          },
			{
				"internalType": "bool",
				"name": "lock",
				"type": "bool"
			}
        ],
        "internalType": "struct Type.Market[]",
        "name": "",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }	
]`
