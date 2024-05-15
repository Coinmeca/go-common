package commonutils

import (
	"encoding/binary"
	"math"
	"math/big"
	"reflect"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/coinmeca/go-common/commonlog"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"fmt"
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
	// Extract the high and low parts of the Decimal128
	high, low := decimal.GetBytes()

	// Determine the sign based on the high part
	isNegative := (high & 0x8000000000000000) != 0

	// If the value is negative, flip all the bits
	if isNegative {
		high = ^high
		low = ^low
	}

	// Combine the high and low parts into a single byte slice
	bytes := make([]byte, 16)
	binary.BigEndian.PutUint64(bytes[:8], high)
	binary.BigEndian.PutUint64(bytes[8:], low)

	// Convert the bytes to a big.Int
	bigInt := new(big.Int).SetBytes(bytes)

	// If the value is negative, negate the big.Int
	if isNegative {
		bigInt.Neg(bigInt)
	}

	return bigInt
}

func Decimal128FromBigInt(bigInt *big.Int) (*primitive.Decimal128, error) {
	if bigInt == nil {
		return nil, errors.New("parameter value is nil")
	}

	// Create a Decimal128 from the string representation of the big.Int
	decimal128, err := primitive.ParseDecimal128(bigInt.String())
	if err != nil {
		return nil, err
	}
	return &decimal128, nil
}

func Decimal128FromFloat64(float float64) (*primitive.Decimal128, error) {
	intValue, frac := math.Modf(float)
	intPart := big.NewInt(int64(intValue))
	fracPart := big.NewInt(int64(frac * math.Pow10(18))) // Assuming 18 decimal places

	var zero float64
	if float < zero {
		intPart = intPart.Neg(intPart)
	}

	decimal128, err := Decimal128FromBigInt(intPart.Add(intPart, fracPart))
	if err != nil {
		return nil, err
	}

	return decimal128, nil
}

func AddDecimal128(decimal1, decimal2 *primitive.Decimal128) (*primitive.Decimal128, error) {
	var zero *primitive.Decimal128
	if decimal1 == zero {
		return decimal2, nil
	} else if decimal2 == zero {
		return decimal1, nil
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(new(big.Int).Add(value1, value2))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SubDecimal128(decimal1, decimal2 *primitive.Decimal128) (*primitive.Decimal128, error) {
	fmt.Println("decimal1",decimal1.String())
	fmt.Println("decimal2",decimal2.String())
	
	var zero *primitive.Decimal128
	if decimal1 == zero {
		return decimal2, nil
	} else if decimal2 == zero {
		return decimal1, nil
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	fmt.Println("value1",value1.String())
	fmt.Println("value2",value2.String())
	fmt.Println("decimal1",decimal1.String())
	fmt.Println("decimal2",decimal2.String())
	fmt.Println("result",new(big.Int).Sub(value1, value2))
	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(new(big.Int).Sub(value1, value2))
	fmt.Println("result",result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// multiply
// value1 4113330394745627368290515537823310743667146809332214747452850902849462206464
// value2 64135250796622190867608690658526625892
// result 4113330394745627368290515537823310743667146809332214747452

func MulDecimal128(decimal1, decimal2 *primitive.Decimal128) (*primitive.Decimal128, error) {
	var zero *primitive.Decimal128
	if decimal1 == zero || decimal2 == zero {
		return zero, nil
	}

	fmt.Println("decimal1", decimal1.String())
	fmt.Println("decimal2", decimal2.String())

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	// Perform multiplication
	value1 = value1.Mul(value1, value2)

	// Convert the result back to primitive.Decimal128
	fmt.Println("multiply")
	fmt.Println("value1", value1.String())
	fmt.Println("value2", value2.String())
	fmt.Println("result", new(big.Int).Div(value1, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)))
	result, err := Decimal128FromBigInt(new(big.Int).Div(value1, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DivDecimal128(decimal1, decimal2 *primitive.Decimal128) (*primitive.Decimal128, error) {
	var zero *primitive.Decimal128
	if decimal1 == zero {
		return decimal2, nil
	} else if decimal2 == zero {
		return decimal1, nil
	}

	value1 := BigIntFromDecimal128(decimal1)
	value1 = new(big.Int).Mul(value1, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	value2 := BigIntFromDecimal128(decimal2)

	fmt.Println("value1",value1.String())
	fmt.Println("value2",value2.String())
	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(new(big.Int).Div(value1, value2))
	fmt.Println("result",result.String())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func QuoDecimal128(decimal1, decimal2 *primitive.Decimal128) (*primitive.Decimal128, error) {
	var zero *primitive.Decimal128
	if decimal1 == zero {
		return decimal2, nil
	} else if decimal2 == zero {
		return decimal1, nil
	}

	value1 := BigIntFromDecimal128(decimal1)
	value1 = new(big.Int).Mul(value1, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	value2 := BigIntFromDecimal128(decimal2)

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(new(big.Int).Quo(value1, value2))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func FloatStringFromDecimal128(decimal *primitive.Decimal128) *string {
	// Convert Decimal128 to string
	str := decimal.String()

	// Add '+' sign if positive
	if str[0] != '-' {
		str = "+" + str
	}

	// Split integer and fractional parts
	parts := strings.SplitN(str, ".", 2)
	integerPart := parts[0]
	fractionalPart := ""
	if len(parts) > 1 {
		fractionalPart = parts[1]
	}

	// Pad fractional part with zeros if needed
	for len(fractionalPart) < 18 {
		fractionalPart += "0"
	}

	// Concatenate integer and fractional parts
	result := integerPart + "." + fractionalPart

	return &result
}
