package mapper

import (
	"encoding/json"
	"fetch-assessment/rules"
)

// assume the raw data for a single item is provided as []byte.
// unmarshal it into an auxiliary struct.
type rawItem struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// implement ItemMapper for raw json
type JsonItemMapper struct{}

func (m *JsonItemMapper) ToDomain(source []byte) (rules.Item, error) {
	var ri rawItem
	if err := json.Unmarshal(source, &ri); err != nil {
		// there was a problem parsing the json for the item, return an empty object and the associated error
		return rules.Item{}, err
	}
	priceCents, err := convertToCents(ri.Price)
	if err != nil {
		// couldn't convert to cents, return an empty object and the associated error
		return rules.Item{}, err
	}
	return rules.Item{
		ShortDescription: ri.ShortDescription,
		Price:            priceCents,
	}, nil
}
