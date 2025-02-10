package validation

import "fetch-assessment/api"

// create a composite to aggregate all the validators
type ReceiptValidationEngine struct {
	Validators []ReceiptValidator
}

/*

this package is for managing validators as defined in the spec

from the yaml:

        Receipt:
            type: object
            required:
                - retailer
                - purchaseDate
                - purchaseTime
                - items
                - total
            properties:
                retailer:
                    description: The name of the retailer or store the receipt is from.
                    type: string
                    pattern: "^[\\w\\s\\-&]+$"
                    example: "M&M Corner Market"
                purchaseDate:
                    description: The date of the purchase printed on the receipt.
                    type: string
                    format: date
                    example: "2022-01-01"
                purchaseTime:
                    description: The time of the purchase printed on the receipt. 24-hour time expected.
                    type: string
                    format: time
                    example: "13:01"
                items:
                    type: array
                    minItems: 1
                    items:
                        $ref: "#/components/schemas/Item"
                total:
                    description: The total amount paid on the receipt.
                    type: string
                    pattern: "^\\d+\\.\\d{2}$"
                    example: "6.49"

*/

func NewReceiptValidationEngine() *ReceiptValidationEngine {
	// validate the retailer
	retailerValidator := NewRetailerReceiptValidator()

	// NOTE: we don't need to validate the purchase date because the server handler that's automatically generated does this for us,
	// and we catch a 500 and return a 400 per the spec anyways, which handles an invalid date (among other issues)

	// validate the purchase time
	purchaseTimeValidator := &PurchaseTimeReceiptValidator{}

	// validate the items
	itemsValidator := &ItemsReceiptValidator{}

	// validate the total
	totalValidator := NewTotalReceiptValidator()

	// aggregate the validators
	return &ReceiptValidationEngine{
		Validators: []ReceiptValidator{
			retailerValidator,
			//purchaseDateValidator // uncessasary
			purchaseTimeValidator,
			itemsValidator,
			totalValidator,
		},
	}
}

// run the validation on the specified receipt
func (e *ReceiptValidationEngine) IsValid(receipt api.Receipt) bool {
	for _, validator := range e.Validators {
		if !validator.IsValid(receipt) {
			return false
		}
	}
	// passed validation for all validators
	return true
}
