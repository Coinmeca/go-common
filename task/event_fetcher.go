package task

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"time"

	ABI "github.com/coinmeca/go-common/abi"
	cv "github.com/coinmeca/go-common/chain"
	rep "github.com/coinmeca/go-common/repository"

	"github.com/coinmeca/go-common/logger"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

func FetchEvents() {
	t := time.Now()

	VaultBlockNo, _ := rep.LastBlockNoFromVaultVolume(CTX)
	if VaultBlockNo == 0 {
		fmt.Println("no last volume block")
	} else {
		VaultBlockNo += 1
	}
	MarketBlockNo, _ := rep.LastBlockNoFromMarketVolume(CTX)
	if MarketBlockNo == 0 {
		fmt.Println("no last market block")
	} else {
		MarketBlockNo += 1
	}

	var blockNo int64
	if VaultBlockNo != 0 && (VaultBlockNo >= MarketBlockNo) {
		blockNo = VaultBlockNo
	} else if MarketBlockNo != 0 && (MarketBlockNo >= VaultBlockNo) {
		blockNo = MarketBlockNo
	} else {
		currentBlock, _ := EthHttpsClient.BlockNumber(context.Background())
		blockNo = int64(currentBlock)
	}

	addresses := append(CAxBOOKS, CAxVAULT)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNo),
		Topics:    [][]common.Hash{{TP1, TP2, TP4, TP5}},
		Addresses: addresses,
	}

	logs, err := EthHttpsClient.FilterLogs(context.Background(), query)
	if err != nil {
		logger.Error("FetchEvents", "err", err, "block number", blockNo)
		panic(err)
	}
	fmt.Printf("Volume logs %v (last block %v)\n", len(logs), blockNo)

	for i, log := range logs {
		fmt.Printf("[%d] volume event: %+v\n", i+1, log)
		switch log.Topics[0] {
		case TP1, TP2:
			setMarketVolume(log)
		case TP4, TP5:
			setVaultVolume(log)
		}
	}

	logger.Info("FetchEvents", "time", time.Since(t).Milliseconds())
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func LogEvents() {
	//transferEventsLog()
	marketEventsLog()
	vaultEventsLog()
}

func marketEventsLog() {
	blockNo, err := rep.LastBlockNoFromMarketVolume(CTX)
	if err != nil {
		fmt.Println("last block error")
		currentBlock, _ := EthHttpsClient.BlockNumber(context.Background())
		blockNo = int64(currentBlock)
	} else {
		blockNo += 1
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(blockNo)),
		Topics:    [][]common.Hash{{TP1, TP2}},
		Addresses: CAxBOOKS,
	}

	logs, err := EthHttpsClient.FilterLogs(context.Background(), query)
	if err != nil {
		logger.Error("marketEventsLog", "err", err)
	}
	fmt.Printf("Market logs %v (last block %v)\n", len(logs), blockNo)

	for i, log := range logs {
		fmt.Printf("[%d] market event: %+v\n", i+1, log)
	}
}

func vaultEventsLog() {
	blockNo, err := rep.LastBlockNoFromVaultVolume(CTX)
	if err != nil {
		fmt.Println("last block error")
		currentBlock, _ := EthHttpsClient.BlockNumber(context.Background())
		blockNo = int64(currentBlock)
	} else {
		blockNo += 1
	}

	currentBlock, _ := EthHttpsClient.BlockNumber(context.Background())
	fmt.Printf("(%v) last block %v\n", CHAINxID, currentBlock)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(currentBlock)),
		Topics:    [][]common.Hash{{TP4, TP5}},
		Addresses: []common.Address{CAxVAULT},
	}

	logs, err := EthHttpsClient.FilterLogs(context.Background(), query)
	if err != nil {
		logger.Error("vaultEventLog", "err", err)
	}
	fmt.Printf("Vault logs %v (last block %v)\n", len(logs), blockNo)

	for i, log := range logs {
		fmt.Printf("[%d] valut event: %+v\n", i+1, log)
	}
}

func transferEventsLog() {
	contractAddress := common.HexToAddress(cv.ETHAddressMumbai)
	topic1, topic2 := common.HexToHash(ABI.ERC20TransferHash), common.HexToHash(ABI.ERC20ApprovalHash)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(33390373), // 10 - 32926529
		//ToBlock:   big.NewInt(33390447),
		Topics:    [][]common.Hash{{topic1, topic2}},
		Addresses: []common.Address{contractAddress},
	}

	logs, err := EthHttpsClient.FilterLogs(context.Background(), query)
	if err != nil {
		logger.Error("transaferEventsLog", "err", err)
	}
	fmt.Printf("logs %v\n", len(logs))

	for _, log := range logs {
		fmt.Printf("BlockNumber: %d", log.BlockNumber)
		//fmt.Printf("BlockHash: %s\n", log.BlockHash)
		//fmt.Printf("TxHash: %s\n", log.TxHash)
		fmt.Printf(" Address: %s", log.Address)
		fmt.Printf(" Event: %s\n", log.Topics[0])
		if len(log.Topics) > 1 {
			fmt.Printf("From: %s", common.HexToAddress(log.Topics[1].String()))
			fmt.Printf(" To: %s\n", common.HexToAddress(log.Topics[2].String()))
		}
		var value int64 = 0
		if len(log.Data) > 0 {
			value, _ = strconv.ParseInt(hex.EncodeToString(log.Data), 16, 64)
			fmt.Printf("Data Value: %v\n", value)
		}
		//rep.SetTransferEventLog(CTX, &log, value)
	}
}
