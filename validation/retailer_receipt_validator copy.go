package validation

import "fetch-assessment/api"

/*
	retailer:
		pattern: "^[\\w\\s\\-&]+$"
*/
func NewRetailerReceiptValidator() *RegexReceiptValidator {
	return NewRegexReceiptValidator(`^[\w\s\-&]+$`, func(receipt api.Receipt) string {
		return receipt.Retailer
	})
}
