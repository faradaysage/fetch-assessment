package validation

import "fetch-assessment/api"

// ensures there is at least one item
type ItemsReceiptValidator struct{}

/*
	items:
		minItems: 1
*/
func (v *ItemsReceiptValidator) IsValid(receipt api.Receipt) bool {
	hasAtLeastOneItem := len(receipt.Items) > 0
	if !hasAtLeastOneItem {
		return false
	}

	// create an item validation engine
	itemValidationEngine := NewItemValidationEngine()

	// validate each item
	for _, item := range receipt.Items {
		if !itemValidationEngine.IsValid(item) {
			return false
		}
	}
	// passed validation for all item validators for all items
	return true
}
