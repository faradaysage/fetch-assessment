package rules

import "time"

// Receipt represents the receipt information.
type Receipt struct {
	Retailer         string
	PurchaseDateTime time.Time
	Total            int64 // represented in cents, alternatively could use something like shopspring/decimal
	Items            []Item
}
