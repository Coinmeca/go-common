package task

import (
	etherchain "coinmeca-go_common/chain"
	"coinmeca-go_common/logger"
	etherrpc "coinmeca-go_common/rpc"
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	WSSClientURI = etherchain.WSSQuickPzkT
)

const IdleDuration = 5 * time.Minute

func NewSubscription(query ethereum.FilterQuery, chEvent chan types.Log) (ethereum.Subscription, error) {
	client, err := etherrpc.NewClient(WSSClientURI)
	if err != nil {
		logger.Error("NewSubscription", "err", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	chSub := make(chan ethereum.Subscription, 1)
	sub, err := client.SubscribeFilterLogs(ctx, query, chEvent)
	if err != nil {
		logger.Error("NewSubscription", "err", err)
		cancel()
	} else {
		chSub <- sub
	}

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout subscription")
	case s := <-chSub:
		return s, nil
	}
}

func SubscribeEventLogs() {
	WSSClientURI = etherchain.WSSProvider[CHAINxID]
	client, err := etherrpc.NewClient(WSSClientURI)
	if err != nil {
		logger.Error("SubscribeEventLogs", "err", err)
		return
	}
	defer client.Close()

	currentBlock, _ := client.BlockNumber(context.Background())
	t := time.Now()
	logger.Info("SubscribeEventLogs", "CHAINxID", CHAINxID, "currentBlock", currentBlock, "now time", t.String())

	addresses := append(CAxBOOKS, CAxVAULT)
	query := ethereum.FilterQuery{
		Addresses: addresses,
		Topics:    [][]common.Hash{{TP1, TP2, TP4, TP5}},
		FromBlock: big.NewInt(int64(currentBlock)),
	}
	eventLogs := make(chan types.Log)
	sub, err := NewSubscription(query, eventLogs)
	if err != nil {
		logger.Error("SubscribeEventLogs", "err", err)
		return
	}

	idleDelay := time.NewTimer(IdleDuration)
	defer idleDelay.Stop()
	for {
		idleDelay.Reset(IdleDuration)
		select {
		case <-sub.Err():
			sub, err = NewSubscription(query, eventLogs)
			if err != nil {
				logger.Error("SubscribeEventLogs", "err", err)
				return
			}
			logger.Info("SubscribeEventLogs", "time", time.Since(t).Seconds())

		case log := <-eventLogs:
			logger.Info("SubscribeEventLogs", "event log", log)
			// do something here with event
			switch log.Topics[0] {
			case TP1, TP2:
				setMarketVolume(log)

			case TP4, TP5:
				setVaultVolume(log)
			}
		case <-idleDelay.C:
			logger.Info("SubscribeEventLogs", "idle delay", time.Since(t).Seconds())
			sub.Unsubscribe()
		}
	}
}
