package mapper

import (
	"math"
	"strconv"
)

// converts a decimal dollar amount, represented by a string to an integer representing the cents portion
// we use an int because Go doesn't have a decimal type, and floating point math isn't precise
// we could use shopspring/decimal or a 3rd party library, but that's unecessary unless a business case is presented for it
func convertToCents(amountStr string) (int64, error) {
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		// if there was an error, default to 0 and bubble up the error so the caller can decide how to handle it
		return 0, err
	}
	// use math.Round to avoid floating-point imprecision
	return int64(math.Round(amount * 100)), nil
}
