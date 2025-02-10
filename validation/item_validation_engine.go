package validation

import "fetch-assessment/api"

// create a composite to aggregate all the validators
type ItemValidationEngine struct {
	Validators []ItemValidator
}

/*

this package is for managing validators as defined in the spec

from the yaml:

        Item:
            required:
                - shortDescription
                - price
            properties:
                shortDescription:
                    pattern: "^[\\w\\s\\-]+$"
                price:
                    pattern: "^\\d+\\.\\d{2}$"

*/

func NewItemValidationEngine() *ItemValidationEngine {
	// validate the short description
	shortDescriptionValidator := NewShortDescriptionItemValidator()

	// validate the price
	priceValidator := NewPriceItemValidator()

	return &ItemValidationEngine{
		Validators: []ItemValidator{
			shortDescriptionValidator,
			priceValidator,
		},
	}
}

// run the validation on the specified item
func (e *ItemValidationEngine) IsValid(item api.Item) bool {
	for _, validator := range e.Validators {
		if !validator.IsValid(item) {
			return false
		}
	}
	// passed validation for all validators
	return true
}
