package commonutils

import (
	"encoding/binary"
	"math/big"

	"github.com/shopspring/decimal"
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
	stringValue := bigInt.String()

	// Create a Decimal128 from the string representation of the big.Int
	decimal128Value, err := primitive.ParseDecimal128(stringValue)
	if err != nil {
		return primitive.Decimal128{}, err
	}

	return decimal128Value, nil
}

func Decimal128FromFloat64(float *float64) primitive.Decimal128 {
	decimal := decimal.NewFromFloat(float)
	high, low := decimalValue.BigParts()
	return primitive.NewDecimal128(high, low)
}
