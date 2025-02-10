package validation

import (
	"fetch-assessment/api"
)

/*

this package is for managing validators as defined in the spec

from the yaml:

    schemas:
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

// define an interface that all our item validators will inherit from
type ItemValidator interface {
	IsValid(receipt api.Item) bool
}
