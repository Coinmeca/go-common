package credit

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

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
				return fmt.Errorf("Failed to set field %s: %v", outputName, err)
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
