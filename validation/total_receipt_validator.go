package validation

import "fetch-assessment/api"

/*
	total:
		pattern: "^\\d+\\.\\d{2}$"
*/
func NewTotalReceiptValidator() *RegexReceiptValidator {
	return NewRegexReceiptValidator(`^\d+\.\d{2}$`, func(receipt api.Receipt) string {
		return receipt.Total
	})
}
