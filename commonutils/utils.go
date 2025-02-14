package commonutils

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

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
	bigInt, _, err := decimal.BigInt()
	if err != nil {
		commonlog.Logger.Warn("BigIntFromDecimal128",
			zap.String("wrong decimal:", bigInt.String()),
		)
		return nil
	}
	return bigInt
}

func Decimal128FromBigInt(bigInt *big.Int) (*primitive.Decimal128, error) {
	// Parse string into Decimal128
	decimal128, err := primitive.ParseDecimal128(bigInt.String())
	if err != nil {
		commonlog.Logger.Warn("Decimal128FromBigInt",
			zap.String("wrong bigInt:", bigInt.String()),
		)
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
			zap.String("nil decimal1", decimal1.String()),
			zap.String("nil decimal2", decimal2.String()),
		)
		return nil
	} else if *decimal1 == zero || decimal1.String() == "0" {
		return decimal2
	} else if *decimal2 == zero || decimal2.String() == "0" {
		return decimal1
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	if value1 == nil || value2 == nil {
		commonlog.Logger.Warn("SubDecimal128",
			zap.String("nil value1", value1.String()),
			zap.String("nil value2", value2.String()),
		)
		return nil
	}

	value2 = new(big.Int).Add(value1, value2)

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(value2)
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
			zap.String("nil decimal1", decimal1.String()),
			zap.String("nil decimal2", decimal2.String()),
		)
		return nil
	} else if (*decimal1 == zero || decimal1.String() == "0") && (*decimal2 == zero || decimal2.String() == "0") {
		return &zero
	} else if *decimal2 == zero || decimal2.String() == "0" {
		return decimal1
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	if value1 == nil || value2 == nil {
		commonlog.Logger.Warn("SubDecimal128",
			zap.String("nil value1", value1.String()),
			zap.String("nil value2", value2.String()),
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
			zap.String("nil decimal1", decimal1.String()),
			zap.String("nil decimal2", decimal2.String()),
		)
		return nil
	} else if *decimal1 == zero || *decimal2 == zero || decimal1.String() == "0" || decimal2.String() == "0" {
		return &zero
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	if value1 == nil || value2 == nil {
		commonlog.Logger.Warn("MulDecimal128",
			zap.String("nil value1", value1.String()),
			zap.String("nil value2", value2.String()),
		)
		return nil
	}

	// Perform multiplication
	value1 = value1.Mul(value1, value2)
	value2 = new(big.Int).Div(value1, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(value2)
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
			zap.String("nil decimal1", decimal1.String()),
			zap.String("nil decimal2", decimal2.String()),
		)
		return nil
	} else if *decimal1 == zero || *decimal2 == zero || decimal1.String() == "0" || decimal2.String() == "0" {
		return &zero
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	if value1 == nil || value2 == nil {
		commonlog.Logger.Warn("SubDecimal128",
			zap.String("nil value1", value1.String()),
			zap.String("nil value2", value2.String()),
		)
		return nil
	}

	value1 = new(big.Int).Mul(value1, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	value2 = new(big.Int).Div(value1, value2)

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(value2)
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
			zap.String("nil decimal1", decimal1.String()),
			zap.String("nil decimal2", decimal2.String()),
		)
		return nil
		// return nil, errors.New("arguments are nil")
	} else if *decimal1 == zero || *decimal2 == zero || decimal1.String() == "0" || decimal2.String() == "0" {
		return &zero
	}

	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

	if value1 == nil || value2 == nil {
		commonlog.Logger.Warn("SubDecimal128",
			zap.String("nil value1", value1.String()),
			zap.String("nil value2", value2.String()),
		)
		return nil
		// return nil, errors.New("nil value")
	}

	value1 = new(big.Int).Mul(value1, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	value2 = new(big.Int).Quo(value1, value2)

	// Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(value2)
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

func FloatStringFromDecimal128(input *primitive.Decimal128) string {
	if input == nil {
		return ""
	}

	value := input.String()

	// // Use a regular expression to remove the decimal point and anything after it
	re := regexp.MustCompile(`\.\d*`)
	value = re.ReplaceAllString(value, "")

	if value == "0" || value == "-0" || value == "0E-6176" || value == "-0E-6176" || strings.TrimLeft(value, "0.-") == "" {
		return "0"
	}

	isNegative := strings.HasPrefix(value, "-")
	if isNegative {
		value = value[1:]
	}

	// Handle special case where the value is a small number in scientific notation
	if strings.Contains(value, "E") {
		if strings.HasPrefix(value, "0E") || strings.HasPrefix(value, "-0E") {
			return "0"
		}
	}

	// Remove leading zeros from the integer part
	numbers := strings.TrimLeft(value, "0")

	if len(numbers) <= 18 {
		// Case where the number has 18 or fewer digits, we treat it as a small number
		zeros := 18 - len(numbers)
		result := fmt.Sprintf("0.%s%s", strings.Repeat("0", zeros), numbers)
		result = strings.TrimRight(result, "0") // Trim trailing zeros
		if strings.HasSuffix(result, ".") {
			result = result[:len(result)-1] // Remove trailing dot if no decimals
		}
		if isNegative {
			return "-" + result
		}
		return result
	}

	// Handle large numbers with more than 18 digits
	// For large numbers, place the decimal point after the 10th digit
	integer := numbers[:len(numbers)-18]
	decimal := numbers[len(numbers)-18:]

	// Combine integer and decimal parts
	result := fmt.Sprintf("%s.%s", integer, decimal)

	// Remove trailing zeros from the decimal part
	result = strings.TrimRight(result, "0")
	if strings.HasSuffix(result, ".") {
		result = result[:len(result)-1] // Remove trailing dot if no decimals
	}

	if isNegative {
		return "-" + result
	}
	return result
}

func IsDecimal128Zero(d primitive.Decimal128) bool {
	str, _, _ := d.BigInt()
	return str.String() == "0"
}

func ConvertUnixToDayStart(unix int) int {
	t := time.Unix(int64(unix), 0)
	return int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC).Unix())
}

func FloatStringToDecimal128(floatStr string) (primitive.Decimal128, error) {
	floatStr = strings.Replace(floatStr, ",", "", -1)

	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		return primitive.Decimal128{}, err
	}

	floatVal = floatVal * math.Pow(10, 18)
	scaledStr := strconv.FormatFloat(floatVal, 'f', -1, 64)

	decimal128, err := primitive.ParseDecimal128(scaledStr)
	if err != nil {
		return primitive.Decimal128{}, err
	}

	return decimal128, nil
}

func ParseInterval(interval string) int64 {
	switch interval {
	case "1":
		return 1
	case "5":
		return 5
	case "15":
		return 15
	case "30":
		return 30
	case "60":
		return 60
	case "120":
		return 120
	case "240":
		return 240
	case "1D":
		return 86400
	case "1W":
		return 604800
	case "1M":
		return 2592000
	default:
		return 1
	}
}

func TruncateUnix(now int64, interval int64) int64 {
	return time.Unix(now, 0).Truncate(time.Duration(interval) * time.Minute).UTC().Unix()
}

func Prettify(value any) string {
	pretty, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return ""
	}
	return string(pretty)
}
