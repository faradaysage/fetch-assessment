package validation

import (
	"fetch-assessment/api"
	"regexp"
)

// pretty much all the receipt validators will inherit from this basic implementation
// since missing items will be initialized with an empty string if they're not passed
// and all fields require at least one character, we don't need to test for edge cases where a field can have an empty string but it's a required field
// basically we don't have any fields that must be explicitly defined as empty rather than implicitly
type RegexReceiptValidator struct {
	Pattern       *regexp.Regexp
	GetStringFunc func(receipt api.Receipt) string
}

// generate a new validator given the specified pattern and a callback to extract the relevant string from the Receipt
func NewRegexReceiptValidator(pattern string, getStringFunc func(receipt api.Receipt) string) *RegexReceiptValidator {
	return &RegexReceiptValidator{
		Pattern:       regexp.MustCompile(pattern),
		GetStringFunc: getStringFunc,
	}
}

// implement the interface
func (r *RegexReceiptValidator) IsValid(receipt api.Receipt) bool {
	// get the relevant string from the receipt
	str := r.GetStringFunc(receipt)
	// return whether it's a match or not
	return r.Pattern.MatchString(str)
}
