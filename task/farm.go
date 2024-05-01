package task

import (
	ABI "github.com/coinmeca/go-common/abi"
	cv "github.com/coinmeca/go-common/chain"
	"github.com/coinmeca/go-common/logger"
	"github.com/coinmeca/go-common/model"
	repo "github.com/coinmeca/go-common/repository"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

func LoadFarmData() {
	dataHex := fmt.Sprintf("0x%s", ABI.FarmSigMap["get_all"])
	contractAddr := common.HexToAddress(cv.MecaAddrMap[CHAINxID]["FARM"])
	callMsg := ethereum.CallMsg{To: &contractAddr, Data: common.FromHex(dataHex)}

	res, err := EthHttpsClient.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		logger.Error("LoadFarmData", "err", err)
		return
	}

	var packList []model.FarmAbiInfo
	err = FarmABI.UnpackIntoInterface(&packList, "get_all", res)
	if err != nil {
		logger.Error("LoadFarmData", "err", err)
		return
	}

	if len(packList) < 1 {
		logger.Debug("LoadFarmData", "data length", len(packList))
		return
	}

	t := time.Now().UTC()
	for _, pack := range packList {
		decimals := TokenDecimals[strings.ToLower(pack.Stake.Hex())]
		farmRow := model.FarmRow{
			Time:        fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", t.Year(), t.Month(), t.Day(), t.Hour(), (t.Minute()/5)*5),
			Address:     strings.ToLower(pack.Farm.Hex()),
			Name:        pack.Name,
			Stake:       strings.ToLower(pack.Stake.Hex()),
			StakeSymbol: pack.StakeSymbol,
			StakeName:   pack.StakeName,
			Earn:        strings.ToLower(pack.Earn.Hex()),
			EarnSymbol:  pack.EarnSymbol,
			EarnName:    pack.EarnName,
			Start:       pack.Start.Uint64(),
			Period:      pack.Period.Uint64(),
			Goal:        pack.Goal.Uint64(),
			Rewards:     decimal.NewFromBigInt(pack.Rewards, -1*decimals),
			Locked:      decimal.NewFromBigInt(pack.Locked, -1*decimals),
		}
		//fmt.Printf("(farm data) %+v\n", farmRow)
		repo.SetFarmData(CTX, farmRow)
	}
}
