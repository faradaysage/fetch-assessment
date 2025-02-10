package validation

import "fetch-assessment/api"

/*
	shortDescription:
		pattern: "^[\\w\\s\\-]+$"
*/
func NewShortDescriptionItemValidator() *RegexItemValidator {
	return NewRegexItemValidator(`^[\w\s\-]+$`, func(item api.Item) string {
		return item.ShortDescription
	})
}
