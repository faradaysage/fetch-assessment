package mapper

import (
	"fetch-assessment/api"
	"testing"
)

func TestApiItemMapper(t *testing.T) {
	// based on morning-receipt.json
	apiItem := api.Item{
		ShortDescription: "Pepsi - 12-oz",
		Price:            "1.25",
	}

	// map the item
	item, err := MapToItem(apiItem)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if item.ShortDescription != "Pepsi - 12-oz" {
		t.Errorf("Expected ShortDescription 'Pepsi - 12-oz', got '%s'", item.ShortDescription)
	}
	if item.Price != 125 {
		t.Errorf("Expected Price '125' cents, got '%d'", item.Price)
	}
}
