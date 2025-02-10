package validation

import (
	"fetch-assessment/api"
)

/*

this package is for managing validators as defined in the spec

from the yaml:

    schemas:
        Receipt:
            required:
                - retailer
                - purchaseDate
                - purchaseTime
                - items
                - total
            properties:
                retailer:
                    pattern: "^[\\w\\s\\-&]+$"
                purchaseDate:
                    format: date
                    example: "2022-01-01"
                purchaseTime:
                    format: time
                    example: "13:01"
                items:
                    minItems: 1
                total:
                    pattern: "^\\d+\\.\\d{2}$"

*/

// define an interface that all our receipt validators will inherit from
type ReceiptValidator interface {
	IsValid(receipt api.Receipt) bool
}
