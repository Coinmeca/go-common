package repository

import (
	"github.com/coinmeca/go-common/logger"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetTransferEventLog(ctx context.Context, pool *pgxpool.Pool, e *types.Log, _value int64) {
	query := `insert into meca.testnet_transfer_logs(block_no,block_hash,tx_hash,address,event,"from","to",value)
values($1,$2,$3,$4,$5,$6,$7,$8) ON CONFLICT ON CONSTRAINT pk_testnet_transfer_logs DO UPDATE SET value = EXCLUDED.value`

	from, to := common.HexToAddress(e.Topics[1].Hex()).Hex(), common.HexToAddress(e.Topics[2].Hex()).Hex()
	_, err := pool.Exec(ctx, query, e.BlockNumber, e.BlockHash.Hex(), e.TxHash.Hex(), e.Address.Hex(), e.Topics[0].Hex(), from, to, _value)
	if err != nil {
		logger.Error("SetTransferEventLog", "err", err, "query", query)
	} else {
		logger.Debug("SetTransferEventLog", "query", query)
	}
}

func SetOtherEventLog(ctx context.Context, pool *pgxpool.Pool, e *types.Log, _value int64) {
	query := `insert into meca.testnet_transfer_logs(block_no,block_hash,tx_hash,address,event,"from","to",value)
values($1,$2,$3,$4,$5,$6,$7,$8) ON CONFLICT ON CONSTRAINT pk_testnet_transfer_logs DO UPDATE SET value = EXCLUDED.value`

	from, to := common.HexToAddress(e.Topics[1].Hex()).Hex(), common.HexToAddress(e.Topics[2].Hex()).Hex()
	_, err := pool.Exec(ctx, query, e.BlockNumber, e.BlockHash.Hex(), e.TxHash.Hex(), e.Address.Hex(), e.Topics[0].Hex(), from, to, _value)
	if err != nil {
		logger.Error("SetOtherEventLog", "err", err, "query", query)
	} else {
		logger.Debug("SetOtherEventLog", "query", query)
	}
}
