package market

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type OutputCheckAccess struct {
	Check bool `abi:""`
}

type OutputFacetAddress struct {
	Address common.Address `abi:""`
}

type OutputFacetAddresses struct {
	Address []common.Address `abi:""`
}

type OutputFacetFunctionSelectors struct {
	Selectors []common.Address `abi:""`
}

type OutputFacets struct {
}

type OutputSupportsInterface struct {
	Interface bool `abi:""`
}

type OutputCheck struct {
	Check bool `abi:""`
}

type OutputLiquidation struct {
	Liquidation bool `abi:""`
}

type OutputCallLimit struct {
	Limit *big.Int `abi:""`
}

type OutputDivider struct {
	Divider *big.Int `abi:""`
}

type OutputFee struct {
	Fee *big.Int `abi:""`
}

type OutputLock struct {
	Lock bool `abi:""`
}

type OutputReward struct {
	Reward *big.Int `abi:""`
}

type OutputThreshold struct {
	Threshold *big.Int `abi:""`
}

type OutputYield struct {
	Yield *big.Int `abi:""`
}

type OutputBase struct {
	Base common.Address `abi:""`
}

type OutputGetInfo struct {
}

type OutputGetOrderBook struct {
}

type OutputGetTicks struct {
}

type OutputNft struct {
	Nft common.Address `abi:""`
}

type OutputPrice struct {
	Price *big.Int `abi:""`
}

type OutputQuote struct {
	Quote common.Address `abi:""`
}

type OutputTick struct {
	Tick *big.Int `abi:""`
}

type OutputGetMargin struct {
	Margin *big.Int `abi:""`
}

type OutputGetThreshold struct {
	Threshold *big.Int `abi:""`
}

type OutputGetApproved struct {
	Approved common.Address `abi:""`
}

type OutputIsApprovedForAll struct {
	All bool `abi:""`
}

type OutputTokenImg struct {
	TokenImg string `abi:""`
}

type OutputBalanceOf struct {
	Balance *big.Int `abi:""`
}

type OutputGetId struct {
	Id *big.Int `abi:"_id"`
}

type OutputGetKey struct {
	Key byte `abi:"_key"`
}

type OutputKeysOf struct {
	Keys [32]byte `abi:"keys"`
}

type OutputOwnerOf struct {
	Owner common.Address `abi:""`
}

type OutputTokensOf struct {
	Tokens []*big.Int `abi:""`
}

type OutputName struct {
	Name string `abi:""`
}

type OutputSymbol struct {
	Symbol string `abi:""`
}

type OutputTotalSupply struct {
	Total *big.Int `abi:""`
}

type OutputTokenUri struct {
	Uri string `abi:""`
}

func Unmarshal(output interface{}, data []byte, contractAbi *abi.ABI, method string) error {
	methodAbi, ok := contractAbi.Methods[method]
	if !ok {
		return fmt.Errorf("method %s not found in ABI", method)
	}

	unpacked, err := methodAbi.Outputs.UnpackValues(data)
	if err != nil {
		return err
	}

	outputValue := reflect.ValueOf(output).Elem()
	outputType := outputValue.Type()

	for i, outputParam := range methodAbi.Outputs {
		if len(unpacked) <= i {
			break
		}

		err := assignField(outputValue, outputType, unpacked[i], outputParam.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func assignField(inputValue reflect.Value, inputType reflect.Type, unpackedValue interface{}, outputName string) error {
	for j := 0; j < inputValue.NumField(); j++ {
		fieldVal := inputValue.Field(j)
		fieldType := inputType.Field(j)

		if fieldType.Tag.Get("abi") == outputName {
			if err := setFieldValue(fieldVal, unpackedValue); err != nil {
				return fmt.Errorf("failed to set field %s: %v", outputName, err)
			}
			break
		}
	}
	return nil
}

func setFieldValue(fieldVal reflect.Value, unpackedValue interface{}) error {
	expectedType := fieldVal.Type()

	switch {
	case expectedType == reflect.TypeOf((*big.Int)(nil)):
		return setBigIntField(fieldVal, unpackedValue)
	case expectedType.Kind() == reflect.String:
		return setStringField(fieldVal, unpackedValue)
	case expectedType.Kind() == reflect.Int:
		return setIntField(fieldVal, unpackedValue)
	case expectedType == reflect.TypeOf(common.Address{}):
		return setAddressField(fieldVal, unpackedValue)
	}

	return fmt.Errorf("unexpected type: %v", expectedType)
}

func setBigIntField(fieldVal reflect.Value, unpackedValue interface{}) error {
	if val, ok := unpackedValue.(*big.Int); ok {
		fieldVal.Set(reflect.ValueOf(val))
		return nil
	}
	return fmt.Errorf("type assertion to *big.Int failed")
}

func setStringField(fieldVal reflect.Value, unpackedValue interface{}) error {
	if val, ok := unpackedValue.(string); ok {
		fieldVal.SetString(val)
		return nil
	}
	return fmt.Errorf("type assertion to string failed")
}

func setIntField(fieldVal reflect.Value, unpackedValue interface{}) error {
	if val, ok := unpackedValue.(*big.Int); ok {
		fieldVal.SetInt(val.Int64())
		return nil
	}
	return fmt.Errorf("type assertion to *big.Int for int field failed")
}

func setAddressField(fieldVal reflect.Value, unpackedValue interface{}) error {
	if val, ok := unpackedValue.(common.Address); ok {
		fieldVal.Set(reflect.ValueOf(val))
		return nil
	}
	return fmt.Errorf("type assertion to common.Address failed")
}
