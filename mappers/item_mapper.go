package mapper

import (
	"fetch-assessment/rules"
)

// a generic interface for mapping a source type T into a rules.Item
type ItemMapper[T any] interface {
	ToDomain(source T) (rules.Item, error)
}
