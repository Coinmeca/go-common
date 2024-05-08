package commonutils

import (
	"encoding/binary"
	"math/big"
	"math"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BigIntFromDecimal128(decimal primitive.Decimal128) *big.Int {
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

func Decimal128FromBigInt(bigInt *big.Int) (primitive.Decimal128, error) {
	// Convert the big.Int to a string
	value := bigInt.String()

	// Create a Decimal128 from the string representation of the big.Int
	decimal128, err := primitive.ParseDecimal128(value)
	if err != nil {
		return primitive.Decimal128{}, err
	}

	return decimal128, nil
}

func Float64ToDecimal128(float float64) (primitive.Decimal128, error) {
    intValue, frac := math.Modf(float)
    intPart := big.NewInt(int64(intValue))
    fracPart := big.NewInt(int64(frac * math.Pow10(18))) // Assuming 18 decimal places

    if float < 0 {
        intPart = intPart.Neg(intPart)
    }

	decimal128, err := Decimal128FromBigInt(intPart.Add(intPart, fracPart))
	if err != nil {
		return primitive.Decimal128{}, err
	}

	return decimal128, nil
}


func MulDecimal128(decimal1, decimal2 primitive.Decimal128) (primitive.Decimal128, error) {
	value1 := BigIntFromDecimal128(decimal1)
	value2 := BigIntFromDecimal128(decimal2)

    // Perform multiplication
    value := new(big.Int).Mul(value1, value2)

    // Convert the result back to primitive.Decimal128
	result, err := Decimal128FromBigInt(value)
	if err != nil {
        return primitive.Decimal128{}, err
	}

    return result, nil
}
