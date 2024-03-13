package commondatabase

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContractInfo struct {
	Id              primitive.ObjectID       `json:"id" bson:"_id"`
	ServiceId       string                   `json:"serviceId" bson:"serviceId"`
	ChainId         string                   `json:"chainId" bson:"chainId"`
	Chain           string                   `json:"chain" bson:"chain"`
	ContractName    string                   `json:"contractName" bson:"contractName"`
	ContractAddress string                   `json:"contractAddress" bson:"contractAddress"`
	Owner           string                   `json:"owner" bson:"owner"`
	Abi             []map[string]interface{} `json:"abi" bson:"abi"`
	Cate            string                   `json:"cate" bson:"cate"`
}

type TxData struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	BlockHash   string             `json:"blockHash" bson:"blockHash"`
	BlockNumber string             `json:"blockNumber" bson:"blockNumber"`
	Hash        string             `json:"hash" bson:"hash"`
	AccessList  []string           `json:"accessList" bson:"accessList"`
	ChainId     string             `json:"chainId" bson:"chainId"`
	From        string             `json:"from" bson:"from"`
	To          string             `json:"to" bson:"to"`
	Gas         string             `json:"gas" bson:"gas"`
	GasPrice    string             `json:"gasPrice" bson:"gasPrice"`
	Input       string             `json:"input" bson:"input"`
	Cate        string             `json:"cate" bson:"cate"`
}

type GrpcTxData struct {
	Id          string `bson:"_id" json:"id"`
	BlockHash   string `json:"blockHash"`
	BlockNumber string `json:"blockNumber"`
	Hash        string `json:"hash"`
	ChainId     string `json:"chainId"`
	From        string `json:"from"`
	To          string `json:"to"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
	Input       string `json:"input"`
	Cate        string `json:"cate"`
}
