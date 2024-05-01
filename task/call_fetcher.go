package task

import (
	"context"
	//ABI "dex-server/internal/abi"
	ABI "github.com/coinmeca/go-common/abi"
	//cv "dex-server/internal/configs"
	cv "github.com/coinmeca/go-common/chain"

	"github.com/ethereum/go-ethereum"

	//"dex-server/internal/logger"
	"github.com/coinmeca/go-common/logger"
	"github.com/coinmeca/go-common/model"

	//"dex-server/internal/model"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func FetchCalls() {
	t := time.Now()

	LoadMarketPrice()
	LoadVaultPrice()
	LoadFarmData()

	logger.Info("FetchCalls", "callList", time.Since(t).Milliseconds())
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CheckCallMsg() {
	t := time.Now()

	vaultCall("get_all")
	vaultCall("get_tokens")
	vaultCall("get_key_tokens")
	farmCall("get_all")

	/*wg := sync.WaitGroup{}
	for _, address := range CAxBOOKS {
		wg.Add(1)
		go func(addr common.Address) {
			marketCall(addr, "get_orderbook")
			marketCall(addr, "get_asks")
			marketCall(addr, "get_bids")
			defer wg.Done()
		}(address)
	}
	wg.Wait()*/

	logger.Debug("CheckCallMsg", "callList", time.Since(t).Milliseconds())
}

func farmCall(sigName string) {
	dataHex := fmt.Sprintf("0x%s", ABI.FarmSigMap[sigName])
	contractAddr := common.HexToAddress(cv.MecaAddrMap[CHAINxID]["FARM"])
	callMsg := ethereum.CallMsg{To: &contractAddr, Data: common.FromHex(dataHex)}

	res, err := EthHttpsClient.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		logger.Error("farmCall", "err", err, "contractAddr", contractAddr)
		return
	}

	var pack []model.FarmAbiInfo
	err = FarmABI.UnpackIntoInterface(&pack, sigName, res)
	if err != nil {
		logger.Error("farmCall", "err", err)
		return
	}

	fmt.Printf("[FARM] unpack data...%+v\n", pack)
}

func marketCall(contractAddr common.Address, sigName string) {
	dataHex := "0x"
	switch sigName {
	case "get_asks", "get_bids", "get_orderbook":
		sig, val := ABI.OrderbookSigMap[sigName], fmt.Sprintf("%064x", 30)
		dataHex = fmt.Sprintf("0x%s%s", sig, val)
	case "get_info", "price", "base", "quote", "tick", "fee":
		dataHex = fmt.Sprintf("0x%s", ABI.OrderbookSigMap[sigName])
	}

	callMsg := ethereum.CallMsg{To: &contractAddr, Data: common.FromHex(dataHex)}

	res, err := EthHttpsClient.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		logger.Error("marketCall", "err", err)
	}

	checkCallResult(res, contractAddr.Hex()[:8], sigName)
}

func vaultCall(name string) {
	dataHex := "0x"
	switch name {
	case "get_asks", "get_bids":
		sig, val := ABI.VaultSigMap[name], fmt.Sprintf("%064x", 10)
		dataHex = fmt.Sprintf("0x%s%s", sig, val)
	case "get_all", "get_tokens", "get_key_tokens":
		dataHex = fmt.Sprintf("0x%s", ABI.VaultSigMap[name])
	}

	contractAddr := common.HexToAddress(cv.MecaAddrMap[CHAINxID]["VAULT"])
	callMsg := ethereum.CallMsg{To: &contractAddr, Data: common.FromHex(dataHex)}

	res, err := EthHttpsClient.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		logger.Error("vaultCall", "err", err)
	}

	var tokens []model.VaultAbiInfo
	err = VaultABI.UnpackIntoInterface(&tokens, name, res)
	if err != nil {
		logger.Error("vaultCall", "err", err)
	}
	fmt.Printf("[VAULT] %s unpack data...%v\n", name, tokens)
}

func checkCallResult(res []byte, orderbookName, sigName string) {
	var err error
	switch sigName {
	case "get_asks", "get_bids":
		var items []struct {
			Price   *big.Int `json:"price"`
			Balance *big.Int `json:"balance"`
		}
		err = OrderbookABI.UnpackIntoInterface(&items, sigName, res)
		if err != nil {
			logger.Error("checkCallResult", "err", err)
		}
		fmt.Printf("(%s) unpack data...%d\n", sigName, len(items))
		for i, item := range items {
			fmt.Printf("[%02d] price: %d, balance: %d \n", i, item.Price, item.Balance)
		}
	case "get_orderbook":
		list, err := OrderbookABI.Unpack(sigName, res)
		if err != nil {
			logger.Error("checkCallResult", "err", err)
		}
		orderList := list[0].(struct {
			Asks []struct {
				Price   *big.Int `json:"price"`
				Balance *big.Int `json:"balance"`
			} `json:"asks"`
			Bids []struct {
				Price   *big.Int `json:"price"`
				Balance *big.Int `json:"balance"`
			} `json:"bids"`
		})
		fmt.Printf("(get_orderbook) unpack data... asks is %v, bids is %v\n", len(orderList.Asks), len(orderList.Bids))
		for i, item := range orderList.Asks {
			fmt.Printf("[ASK:%02d] price: %d, balance: %d \n", i+1, item.Price, item.Balance)
		}
		for i, item := range orderList.Bids {
			fmt.Printf("[BID:%02d] price: %d, balance: %d \n", i+1, item.Price, item.Balance)
		}
	case "price", "base", "quote", "tick", "fee":
		val, err := OrderbookABI.Unpack(sigName, res)
		if err != nil {
			logger.Error("checkCallResult", "err", err)
		}
		fmt.Printf("[%s] (%s) unpack data...%v \n", orderbookName, sigName, val[0])
	}
}
