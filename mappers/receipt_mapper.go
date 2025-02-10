package mapper

import (
	"fetch-assessment/rules"
)

// a generic interface for mapping a source type T into a rules.Receipt
type ReceiptMapper[T any] interface {
	ToDomain(source T) (rules.Receipt, error)
}
