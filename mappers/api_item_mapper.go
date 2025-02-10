package mapper

import (
	"fetch-assessment/api"
	"fetch-assessment/rules"
)

// implement ItemMapper for api.Item
type ApiItemMapper struct{}

func (m *ApiItemMapper) ToDomain(source api.Item) (rules.Item, error) {
	priceCents, err := convertToCents(source.Price)
	if err != nil {
		// couldn't convert to cents, return an empty object and the associated error
		return rules.Item{}, err
	}
	return rules.Item{
		ShortDescription: source.ShortDescription,
		Price:            priceCents,
	}, nil
}
