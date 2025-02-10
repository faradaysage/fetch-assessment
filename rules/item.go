package rules

// Item represents the information for each item in a receipt.
type Item struct {
	ShortDescription string
	Price            int64 // represented in cents, alternatively could use something like shopspring/decimal
}
