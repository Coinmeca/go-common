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

type BatchInfo struct {
	Id              primitive.ObjectID `json:"id" bson:"_id"`
	BatchId         int                `json:"batchId" bson:"batchId"`
	Term            string             `json:"term" bson:"term"`
	Cate            string             `json:"cate" bson:"cate"`
	Tilockede           string             `json:"tilockede" bson:"tilockede"`
	ContractAddress string             `json:"contractAddress" bson:"contractAddress"`
	Chain           string             `json:"chain" bson:"chain"`
	Description     string             `json:"description" bson:"description"`
}

type ChainInfo struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ChainId   int32              `json:"chainId" bson:"chainId"`
	ChainName string             `json:"chainName" bson:"chainName"`
	Currency  string             `json:"currency" bson:"currency"`
	Base      string             `json:"base" bson:"base"`
	Type      string             `json:"type" bson:"type"`
	Rpc       string             `json:"rpc" bson:"rpc"`
}
