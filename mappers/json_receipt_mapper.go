package mapper

import (
	"encoding/json"
	"fetch-assessment/rules"
	"time"
)

// assume the raw receipt is provided as []byte.
// unmarshal into an auxiliary structure.
type rawReceipt struct {
	Retailer     string            `json:"retailer"`
	PurchaseDate string            `json:"purchaseDate"`
	PurchaseTime string            `json:"purchaseTime"`
	Total        string            `json:"total"`
	Items        []json.RawMessage `json:"items"` // each item as raw json
}

// implements ReceiptMapper for api.Receipt.
type JsonReceiptMapper struct {
	// inject the item mapper
	ItemMapper ItemMapper[[]byte]
}

func (m *JsonReceiptMapper) ToDomain(source []byte) (rules.Receipt, error) {
	var rr rawReceipt
	if err := json.Unmarshal(source, &rr); err != nil {
		return rules.Receipt{}, err
	}

	combined := rr.PurchaseDate + " " + rr.PurchaseTime
	const layout = "2006-01-02 15:04" // 01/02 03:04:05PM ‘06 -0700, https://yourbasic.org/golang/format-parse-string-time-date-example/
	parsedTime, err := time.Parse(layout, combined)
	if err != nil {
		// could not parse the time, return an empty object and bubble the error
		return rules.Receipt{}, err
	}

	totalCents, err := convertToCents(rr.Total)
	if err != nil {
		// could not parse the dollar (cents), return an empty object and bubble up the error
		return rules.Receipt{}, err
	}

	// map each item using the injected ItemMapper
	var domainItems []rules.Item
	for _, rawItemData := range rr.Items {
		item, err := m.ItemMapper.ToDomain(rawItemData)
		if err != nil {
			// problem mapping the item, return an empty Receipt along with the error
			return rules.Receipt{}, err
		}
		domainItems = append(domainItems, item)
	}

	return rules.Receipt{
		Retailer:         rr.Retailer,
		PurchaseDateTime: parsedTime,
		Total:            totalCents,
		Items:            domainItems,
	}, nil
}
