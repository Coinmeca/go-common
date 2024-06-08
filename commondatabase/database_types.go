package commondatabase

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contract struct {
	Id        primitive.ObjectID        `json:"id" bson:"_id"`
	ServiceId string                    `json:"serviceId" bson:"serviceId"`
	Cate      string                    `json:"cate" bson:"cate"`
	ChainId   string                    `json:"chainId" bson:"chainId"`
	Name      string                    `json:"name" bson:"name"`
	Address   string                    `json:"address" bson:"address"`
	Owner     string                    `json:"owner" bson:"owner"`
	Abi       *[]map[string]interface{} `json:"abi" bson:"abi"`
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
	BatchId         int32              `json:"batchId" bson:"batchId"`
	Term            string             `json:"term" bson:"term"`
	Cate            string             `json:"cate" bson:"cate"`
	Title           string             `json:"title" bson:"title"`
	ContractAddress string             `json:"contractAddress" bson:"contractAddress"`
	Chain           string             `json:"chain" bson:"chain"`
	ChainId         string             `json:"chainId" bson:"chainId"`
	Description     string             `json:"description" bson:"description"`
}

type Currency struct {
	Name     string `json:"name" bson:"name"`
	Symbol   string `json:"symbol" bson:"symbol"`
	Decimals int    `json:"decimals" bson:"decimals"`
}

type Chain struct {
	Id             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ChainId        string             `json:"chainId" bson:"chainId"`
	ChainName      string             `json:"chainName" bson:"chainName"`
	NativeCurrency Currency           `json:"nativeCurrency" bson:"nativeCurrency"`
	Base           string             `json:"base" bson:"base"`
	Type           string             `json:"type" bson:"type"`
	Rpc            string             `json:"rpc" bson:"rpc"`
}

type APIKey struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Url       string             `json:"url" bson:"url"`
	Key       string             `json:"key" bson:"key"`
	Active    bool               `json:"active" bson:"active"`
	Start     int                `json:"start" bson:"start,omitempty"`
	Expired   int                `json:"expired" bson:"expired,omitempty"`
	Retry     int                `json:"retry" bson:"retry"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}
