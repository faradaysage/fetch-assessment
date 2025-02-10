package validation

import "fetch-assessment/api"

/*
	price:
		pattern: "^\\d+\\.\\d{2}$"
*/
func NewPriceItemValidator() *RegexItemValidator {
	return NewRegexItemValidator(`^\d+\.\d{2}$`, func(item api.Item) string {
		return item.Price
	})
}
