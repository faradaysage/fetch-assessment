package mapper

import (
	"testing"
)

// just a simple test to make sure the json item mapper actually works as expected
// we could get more extensive, testing a variety of scenarios, but this creates a baseline
func TestJsonItemMapperSimple(t *testing.T) {
	// get a sample of an expected json item (this was grabbed from morning-receipt.json)
	itemJson := []byte(`{"shortDescription": "Pepsi - 12-oz", "price": "1.25"}`)

	// map the item
	item, err := MapToItem(itemJson)
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
