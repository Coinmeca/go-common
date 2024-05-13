package commonutils

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"reflect"
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
	// Extract the low part of the Decimal128
	_, low := decimal.GetBytes()

	// Encode the low part into a byte slice
	lowBytes := make([]byte, binary.MaxVarintLen64)
	numBytes := binary.PutUvarint(lowBytes, uint64(low))
	lowBytes = lowBytes[:numBytes]

	// Combine the low part into a BigInt
	bigIntValue := new(big.Int)
	bigIntValue.SetBytes(lowBytes)

	return bigIntValue
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
	zero := new(big.Int)
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
	zero := new(big.Int)
	if decimal1 == zero {
		return decimal2, nil
	} else if decimal2 == zero {
		return decimal1, nil
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(new(big.Int).Sub(value1, value2))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func MulDecimal128(decimal1, decimal2 *primitive.Decimal128) (*primitive.Decimal128, error) {
	var zero *primitive.Decimal128
	if decimal1 == zero || decimal2 == zero {
		return zero, nil
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	// Perform multiplication
	value1 = value1.Mul(value1, value2)

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(new(big.Int).Div(value1, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)))
	if err != nil {
		return nil, err
	}

	fmt.Println("decimal value: ", result)
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

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(new(big.Int).Div(value1, value2))
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
