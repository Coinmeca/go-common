package vault

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"reflect"
	"strings"
)

const (
	MethodGetAll       = "getAll"
	MethodGetKeyTokens = "getKeyTokens"
)

const (
	TradeTypeDeposit = iota
	TradeTypeWithdraw
)

type TokenInfo struct {
	Key      bool           `abi:"key"`
	Address  common.Address `abi:"addr"`
	Name     string         `abi:"name"`
	Symbol   string         `abi:"symbol"`
	Decimals uint8          `abi:"decimals"`
	Treasury *big.Int       `abi:"treasury"`
	Rate     *big.Int       `abi:"rate"`
	Weight   *big.Int       `abi:"weight"`
	Need     *big.Int       `abi:"need"`
}

type OutputGetAll struct {
	Tokens []TokenInfo `abi:""`
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

func UnmarshalV3(abiData *abi.ABI, methodName string, data []byte, output interface{}) error {
	method, exists := abiData.Methods[methodName]
	if !exists {
		return fmt.Errorf("%s method not found in ABI", methodName)
	}

	rawOutput, err := method.Outputs.Unpack(data)
	if err != nil {
		return fmt.Errorf("failed to unpack output: %v", err)
	}

	fmt.Printf("Raw output: %+v\n", rawOutput)

	val := reflect.ValueOf(output).Elem()
	if len(rawOutput) != val.NumField() {
		return fmt.Errorf("field count mismatch:expected %d, got %d", val.NumField(), len(rawOutput))
	}

	for i, rawVal := range rawOutput {
		field := val.Field(i)
		if !field.CanSet() {
			continue
		}

		fieldValue := reflect.ValueOf(rawVal)
		if field.Type() != fieldValue.Type() {
			return fmt.Errorf("type mismatch for field %d: expected %s, got %s", i, field.Type(), fieldValue.Type())
		}

		field.Set(fieldValue)
	}

	return nil
}
func UnmarshalABI(contractABI *abi.ABI, methodName string, data []byte, output interface{}) error {
	method, exists := contractABI.Methods[methodName]
	if !exists {
		return fmt.Errorf("method %s not found in ABI", methodName)
	}

	unpackedData, err := method.Outputs.Unpack(data)
	if err != nil {
		return fmt.Errorf("unpacking data failed: %v", err)
	}

	if err := mapToStruct(method.Outputs, unpackedData, output); err != nil {
		return fmt.Errorf("mapping unpacked data to struct failed: %v", err)
	}

	return nil
}

func mapToStruct(outputs abi.Arguments, unpackedData []interface{}, output interface{}) error {
	val := reflect.ValueOf(output).Elem()
	if !val.IsValid() {
		return fmt.Errorf("invalid output interface")
	}

	for i, arg := range outputs {
		fieldName := strings.Title(arg.Name)
		structField := val.FieldByName(fieldName)
		if !structField.IsValid() {
			fmt.Printf("Field %s not found in output struct\n", fieldName)
			continue
		}

		fmt.Printf("Mapping field: %s, Value: %v\n", fieldName, unpackedData[i])
		err := setField(structField, unpackedData[i])
		if err != nil {
			return fmt.Errorf("setting field %s failed: %v", fieldName, err)
		}
	}
	return nil
}

func setField(field reflect.Value, value interface{}) error {
	if !field.CanSet() {
		return fmt.Errorf("cannot set field: %v", field.Type())
	}

	fieldValue := reflect.ValueOf(value)
	if fieldValue.Type().AssignableTo(field.Type()) {
		field.Set(fieldValue)
	} else if fieldValue.Type().ConvertibleTo(field.Type()) {
		field.Set(fieldValue.Convert(field.Type()))
	} else {
		return fmt.Errorf("type mismatch: cannot convert or assign %v to %v", fieldValue.Type(), field.Type())
	}

	return nil
}

func UnmarshalV2(data []byte, contractAbi *abi.ABI, methodName string, output interface{}) error {
	method, exists := contractAbi.Methods[methodName]
	if !exists {
		return fmt.Errorf("%s method not found", methodName)
	}

	values, err := method.Outputs.Unpack(data)
	if err != nil {
		return err
	}

	outputValue := reflect.ValueOf(output).Elem()
	if len(values) != outputValue.NumField() {
		return fmt.Errorf("expected %d values, got %d", outputValue.NumField(), len(values))
	}

	for i, value := range values {
		fieldValue := outputValue.Field(i)

		if !fieldValue.Type().AssignableTo(reflect.TypeOf(value)) {
			return fmt.Errorf("type mismatch for field %d", i)
		}

		fieldValue.Set(reflect.ValueOf(value))
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
