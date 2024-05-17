package commonutils

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/coinmeca/go-common/commonlog"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func JoinFromStructs(slice interface{}, fieldName, sep string) string {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		panic("JoinStructs: slice argument must be a slice")
	}

	var parts []string
	for i := 0; i < sliceValue.Len(); i++ {
		elem := sliceValue.Index(i)
		if elem.Kind() == reflect.Struct {
			field := elem.FieldByName(fieldName)
			if !field.IsValid() {
				commonlog.Logger.Error("JoinFromStructs",
					zap.String("field '%s' not found in struct", fieldName),
				)
				continue
			}

			if field.Kind() != reflect.String {
				commonlog.Logger.Error("JoinFromStructs",
					zap.String("field '%s' must be of type string", fieldName),
				)
				continue
			}
			parts = append(parts, field.String())
		} else {
			commonlog.Logger.Error("JoinFromStructs",
				zap.String("Not struct type", ""),
			)
			continue
		}
	}

	return strings.Join(parts, sep)
}

func GetCurrentDate() *string {
	currentTime := time.Now()
	formattedDate := currentTime.Format("2006-01-02 15:04:05")
	return &formattedDate
}

func FormattedDate(t *int64) *string {
	unixTime := time.Unix(*t, 0)
	formattedDate := unixTime.Format("2006-01-02 15:04:05")
	return &formattedDate
}

func BigIntFromDecimal128(decimal *primitive.Decimal128) *big.Int {
	// Attempt to set string value to big.Int
	value := new(big.Int)
	if decimal == nil {
		fmt.Println("[BigIntFromDecimal128] wrong decimal:", decimal)
		return nil
	}
	if *decimal == primitive.NewDecimal128(0, 0) {
		return value
	}
	value.SetString(decimal.String(), 10)

	return value
}

func Decimal128FromBigInt(bigInt *big.Int) (*primitive.Decimal128, error) {
	if bigInt == nil {
		fmt.Println("[Decimal128FromBigInt] wrong bigInt:", bigInt)
		return nil, errors.New("bigInt is nil")
	}

	// Create a Decimal128 from the string representation of the big.Int
	decimal128, err := primitive.ParseDecimal128(bigInt.String())
	if err != nil {
		fmt.Println("[Decimal128FromBigInt] failed to parse:", decimal128)
		return nil, err
	}
	return &decimal128, nil
}

func Decimal128FromFloat64(float float64) (*primitive.Decimal128, error) {
	floatString := strconv.FormatFloat(float*math.Pow(10, 18), 'f', -1, 64)

	// Parse the string to Decimal128
	decimal128, err := primitive.ParseDecimal128(floatString)
	if err != nil {
		return &primitive.Decimal128{}, fmt.Errorf("error parsing Decimal128: %v", err)
	}

	return &decimal128, nil
}

func AddDecimal128(decimal1, decimal2 *primitive.Decimal128) *primitive.Decimal128 {
	var zero primitive.Decimal128
	if decimal1 == nil || decimal2 == nil {
		commonlog.Logger.Warn("AddDecimal128",
			zap.String("decimal1", decimal1.String()),
			zap.String("decimal2", decimal2.String()),
		)
		return nil
	} else if *decimal1 == zero {
		return decimal2
	} else if *decimal2 == zero {
		return decimal1
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	if value1 == nil || value2 == nil {
		commonlog.Logger.Warn("SubDecimal128",
			zap.String("value1", value1.String()),
			zap.String("value2", value2.String()),
		)
		return nil
	}

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(new(big.Int).Add(value1, value2))
	if err != nil {
		commonlog.Logger.Warn("AddDecimal128",
			zap.String("value1", value1.String()),
			zap.String("value2", value2.String()),
			zap.String("result", result.String()),
		)
		return nil
	}

	return result
}

func SubDecimal128(decimal1, decimal2 *primitive.Decimal128) *primitive.Decimal128 {
	var zero primitive.Decimal128
	if decimal1 == nil || decimal2 == nil {
		commonlog.Logger.Warn("SubDecimal128",
			zap.String("decimal1", decimal1.String()),
			zap.String("decimal2", decimal2.String()),
		)
		return nil
	} else if *decimal1 == zero && *decimal2 == zero {
		return &zero
	} else if *decimal2 == zero {
		return decimal1
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	if value1 == nil || value2 == nil {
		commonlog.Logger.Warn("SubDecimal128",
			zap.String("value1", value1.String()),
			zap.String("value2", value2.String()),
		)
		return nil
	}

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(new(big.Int).Sub(value1, value2))
	if err != nil {
		commonlog.Logger.Warn("SubDecimal128",
			zap.String("value1", value1.String()),
			zap.String("value2", value2.String()),
			zap.String("result", result.String()),
		)
		return nil
	}

	return result
}

func MulDecimal128(decimal1, decimal2 *primitive.Decimal128) *primitive.Decimal128 {
	var zero primitive.Decimal128
	if decimal1 == nil || decimal2 == nil {
		commonlog.Logger.Warn("MulDecimal128",
			zap.String("decimal1", decimal1.String()),
			zap.String("decimal2", decimal2.String()),
		)
		return nil
	} else if *decimal1 == zero || *decimal2 == zero {
		return &zero
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	if value1 == nil || value2 == nil {
		commonlog.Logger.Warn("SubDecimal128",
			zap.String("value1", value1.String()),
			zap.String("value2", value2.String()),
		)
		return nil
	}

	// Perform multiplication
	value1 = value1.Mul(value1, value2)

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(new(big.Int).Div(value1, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)))
	if err != nil {
		commonlog.Logger.Warn("MulDecimal128",
			zap.String("value1", value1.String()),
			zap.String("value2", value2.String()),
			zap.String("result", result.String()),
		)
		return nil
	}

	return result
}

func DivDecimal128(decimal1, decimal2 *primitive.Decimal128) *primitive.Decimal128 {
	var zero primitive.Decimal128
	if decimal1 == nil || decimal2 == nil {
		commonlog.Logger.Warn("DivDecimal128",
			zap.String("decimal1", decimal1.String()),
			zap.String("decimal2", decimal2.String()),
		)
		return nil
	} else if *decimal1 == zero || *decimal2 == zero {
		return &zero
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	if value1 == nil || value2 == nil {
		commonlog.Logger.Warn("SubDecimal128",
			zap.String("value1", value1.String()),
			zap.String("value2", value2.String()),
		)
		return nil
	}

	value1 = new(big.Int).Mul(value1, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(new(big.Int).Div(value1, value2))
	if err != nil {
		commonlog.Logger.Warn("DivDecimal128",
			zap.String("value1", value1.String()),
			zap.String("value2", value2.String()),
			zap.String("result", result.String()),
		)
		return nil
	}

	return result
}

func QuoDecimal128(decimal1, decimal2 *primitive.Decimal128) *primitive.Decimal128 {
	var zero primitive.Decimal128
	if decimal1 == nil || decimal2 == nil {
		commonlog.Logger.Warn("QuoDecimal128",
			zap.String("decimal1", decimal1.String()),
			zap.String("decimal2", decimal2.String()),
		)
		return nil
		// return nil, errors.New("arguments are nil")
	} else if *decimal1 == zero || *decimal2 == zero {
		return &zero
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	if value1 == nil || value2 == nil {
		commonlog.Logger.Warn("SubDecimal128",
			zap.String("value1", value1.String()),
			zap.String("value2", value2.String()),
		)
		return nil
		// return nil, errors.New("nil value")
	}

	value1 = new(big.Int).Mul(value1, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(new(big.Int).Quo(value1, value2))
	if err != nil {
		commonlog.Logger.Warn("QuoDecimal128",
			zap.String("value1", value1.String()),
			zap.String("value2", value2.String()),
			zap.String("result", result.String()),
		)
		return nil
	}

	return result
}

func FloatStringFromDecimal128(decimal *primitive.Decimal128) string {
	// Convert the big.Int to a string
	var result string
	if decimal == nil {
		return ""
	}

	var zero primitive.Decimal128
	if *decimal == zero {
		return "0"
	}

	bigInt, _ := new(big.Int).SetString(decimal.String(), 10)
	value := bigInt.String()

	if value == "0" {
		return value
	}

	// Determine the length of the string
	length := len(value)

	// Check if the length is less than the number of decimal places
	if length <= 18 {
		// Pad the string with leading zeros if necessary
		prefix := "0."
		suffix := value
		for i := 0; i < 18-length; i++ {
			prefix += "0"
		}
		result = prefix + suffix
		return result
	}

	// Insert the decimal point at the appropriate position
	index := length - 18
	integer := value[:index]
	float := value[index:]

	// Remove trailing zeros from the decimal part
	float = strings.TrimRight(float, "0")

	// If the decimal part is empty after removing trailing zeros,
	// return only the integer part
	if float == "" {
		return integer
	}

	// Pad the decimal part with trailing zeros if necessary
	for len(float) < 18 {
		float += "0"
	}

	// Remove any trailing decimal point
	if float[len(float)-1] == '.' {
		float = float[:len(float)-1]
	}

	result = integer + "." + float
	return result
}
