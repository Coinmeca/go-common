package commonprotocol

const (
	TxStatusSuccess = 1
	TxStatusFailed  = 2
)

type TransactionRecord struct {
	Hash                 string `bson:"hash"`
	Status               uint64 `bson:"status"`
	BlockNumber          uint64 `bson:"blockNumber"`
	Timestamp            uint64 `bson:"timestamp"`
	From                 string `bson:"from"`
	To                   string `bson:"to,omitempty"`
	Value                string `bson:"value"`
	GasPrice             string `bson:"gasPrice,omitempty"`
	MaxPriorityFeePerGas string `bson:"maxPriorityFeePerGas,omitempty"`
	MaxFeePerGas         string `bson:"maxFeePerGas,omitempty"`
	GasUsed              uint64 `bson:"gasUsed"`
	GasLimit             uint64 `bson:"gasLimit"`
	Nonce                uint64 `bson:"nonce"`
	TxType               uint8  `bson:"txType"`
	InputData            string `bson:"inputData"`
}
