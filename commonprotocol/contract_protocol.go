package commonprotocol

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contract struct {
	Id        primitive.ObjectID
	ServiceId string
	Cate      string
	Name      string
	Address   string
	Owner     string
	Abi       *abi.ABI
}

type EthBlock struct {
	BaseFeePerGas    string        `json:"baseFeePerGas"`
	ExtraData        string        `json:"extraData"`
	GasLimit         string        `json:"gasLimit"`
	GasUsed          string        `json:"gasUsed"`
	LogsBloom        string        `json:"logsBloom"`
	Miner            string        `json:"miner"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	Timestamp        string        `json:"timestamp"`
	Hash             string        `json:"hash"`
	MixHash          string        `json:"mixHash"`
	Number           string        `json:"number"`
	ParentHash       string        `json:"parentHash"`
	TotalDifficulty  string        `json:"totalDifficulty"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []interface{} `json:"uncles"`
	Difficulty       string        `json:"difficulty"`
	Nonce            string        `json:"nonce"`
	Sha3Uncles       string        `json:"sha3Uncles"`
	Size             string        `json:"size"`
	StateRoot        string        `json:"stateRoot"`
	Transactions     []interface{} `json:"transactions"`
}

type BlockData struct {
	Difficulty       string        `json:"difficulty"`
	ExtraData        string        `json:"extraData"`
	GasLimit         string        `json:"gasLimit"`
	GasUsed          string        `json:"gasUsed"`
	Hash             string        `json:"hash"`
	LogsBloom        string        `json:"logsBloom"`
	Miner            string        `json:"miner"`
	MixHash          string        `json:"mixHash"`
	Nonce            string        `json:"nonce"`
	Number           string        `json:"number"`
	ParentHash       string        `json:"parentHash"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	Sha3Uncles       string        `json:"sha3Uncles"`
	Size             string        `json:"size"`
	StateRoot        string        `json:"stateRoot"`
	Timestamp        string        `json:"timestamp"`
	Transactions     []string      `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []common.Hash `json:"uncles"`
}

type Log struct {
	TransactionHash string `json:"transactionHash"`
	Address         string `json:"address"`
	BlockHash       string `json:"blockHash"`
	BlockNumber     string `json:"blockNumber"`
	Data            string `json:"data"`
	//Address     common.Address `json:"address" gencodec:"required"`
	//Topics      []common.Hash  `json:"topics" gencodec:"required"`
	//Data        []byte         `json:"data" gencodec:"required"`
	//BlockNumber uint64         `json:"blockNumber" rlp:"-"`
	//TxHash      common.Hash    `json:"transactionHash" gencodec:"required" rlp:"-"`
	//TxIndex     uint           `json:"transactionIndex" rlp:"-"`
	//BlockHash   common.Hash    `json:"blockHash" rlp:"-"`
	//Index       uint           `json:"logIndex" rlp:"-"`
	//Removed     bool           `json:"removed" rlp:"-"`
}

type ReceiptData struct {
	TransactionHash   string `json:"transactionHash"`
	BlockHash         string `json:"blockHash"`
	BlockNumber       string `json:"blockNumber"`
	Logs              []*Log `json:"logs"`
	ContractAddress   string `json:"contractAddress"`
	EffectiveGasPrice string `json:"effectiveGasPrice"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	From              string `json:"from"`
	GasUsed           string `json:"gasUsed"`
	LogsBloom         string `json:"logsBloom"`

	//Type              uint8          `json:"type,omitempty"`
	//PostState         []byte         `json:"root"`
	//Status            uint64         `json:"status"`
	//CumulativeGasUsed uint64         `json:"cumulativeGasUsed" gencodec:"required"`
	//Bloom             Bloom          `json:"logsBloom"         gencodec:"required"`
	//Logs              []*Log         `json:"logs"              gencodec:"required"`
	//TxHash            common.Hash    `json:"transactionHash" gencodec:"required"`
	//ContractAddress   common.Address `json:"contractAddress"`
	//GasUsed           uint64         `json:"gasUsed" gencodec:"required"`
	//EffectiveGasPrice *big.Int       `json:"effectiveGasPrice"`
	//BlobGasUsed       uint64         `json:"blobGasUsed,omitempty"`
	//BlobGasPrice      *big.Int       `json:"blobGasPrice,omitempty"`
	//BlockHash         common.Hash    `json:"blockHash,omitempty"`
	//BlockNumber       *big.Int       `json:"blockNumber,omitempty"`
	//TransactionIndex  uint           `json:"transactionIndex"`
}

type ContractEventManage struct {
	Address     string         `json:"address" bson:"address"`
	EventTarget *[]EventTarget `json:"target" bson:"target"`
	Cate        string         `json:"cate" bson:"cate"`
}

type EventTarget struct {
	Topic     string `json:"topic" bson:"topic"`
	TopicHash string `json:"topicHash" bson:"topicHash"`
	Address   string `json:"address" bson:"address"`
	Func      string `json:"func" bson:"func"`
	Symbol    string `json:"symbol" bson:"symbol"`
}
