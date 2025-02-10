package mapper

import (
	"fetch-assessment/api"
	"fetch-assessment/rules"
	"time"
)

// implements ReceiptMapper for api.Receipt.
type ApiReceiptMapper struct {
	// inject the item mapper
	ItemMapper ItemMapper[api.Item]
}

func (m *ApiReceiptMapper) ToDomain(source api.Receipt) (rules.Receipt, error) {
	combined := source.PurchaseDate.String() + " " + source.PurchaseTime
	const layout = "2006-01-02 15:04" // 01/02 03:04:05PM ‘06 -0700, https://yourbasic.org/golang/format-parse-string-time-date-example/
	parsedTime, err := time.Parse(layout, combined)
	if err != nil {
		// could not parse the time, return an empty object and bubble the error
		return rules.Receipt{}, err
	}

	totalCents, err := convertToCents(source.Total)
	if err != nil {
		// could not parse the dollar (cents), return an empty object and bubble up the error
		return rules.Receipt{}, err
	}

	// map each item using the injected ItemMapper
	var domainItems []rules.Item
	for _, apiItem := range source.Items {
		item, err := m.ItemMapper.ToDomain(apiItem)
		if err != nil {
			// problem mapping the item, return an empty Receipt along with the error
			return rules.Receipt{}, err
		}
		// add the item
		domainItems = append(domainItems, item)
	}

	return rules.Receipt{
		Retailer:         source.Retailer,
		PurchaseDateTime: parsedTime,
		Total:            totalCents,
		Items:            domainItems,
	}, nil
}
