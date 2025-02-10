package validation

import (
	"fetch-assessment/api"
	"regexp"
)

// pretty much all the item validators will inherit from this basic implementation
// since missing items will be initialized with an empty string if they're not passed
// and all fields require at least one character, we don't need to test for edge cases where a field can have an empty string but it's a required field
// basically we don't have any fields that must be explicitly defined as empty rather than implicitly
type RegexItemValidator struct {
	Pattern       *regexp.Regexp
	GetStringFunc func(item api.Item) string
}

// generate a new validator given the specified pattern and a callback to extract the relevant string from the item
func NewRegexItemValidator(pattern string, getStringFunc func(item api.Item) string) *RegexItemValidator {
	return &RegexItemValidator{
		Pattern:       regexp.MustCompile(pattern),
		GetStringFunc: getStringFunc,
	}
}

// implement the interface
func (r *RegexItemValidator) IsValid(item api.Item) bool {
	// get the relevant string from the item
	str := r.GetStringFunc(item)
	// return whether it's a match or not
	return r.Pattern.MatchString(str)
}
