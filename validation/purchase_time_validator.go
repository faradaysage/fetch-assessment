package validation

import (
	"fetch-assessment/api"
	"time"
)

// ensures the time is a valid 24-hour format
type PurchaseTimeReceiptValidator struct{}

/*
purchaseTime:

	format: time
	example: "13:01"
*/
func (v *PurchaseTimeReceiptValidator) IsValid(receipt api.Receipt) bool {
	_, err := time.Parse("15:04", receipt.PurchaseTime)
	return err == nil
}
