package mapper

import (
	"fetch-assessment/api"
	"fetch-assessment/rules"
	"testing"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func TestApiReceiptMapper(t *testing.T) {
	// create a datetime object
	purchaseDate, _ := time.Parse("2006-01-02", "2022-01-02")

	// based on morning-receipt.json
	apiReceipt := api.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: openapi_types.Date{Time: purchaseDate},
		PurchaseTime: "08:13",
		Total:        "2.65",
		Items: []api.Item{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "Dasani", Price: "1.40"},
		},
	}

	// map to our business object
	receipt, err := MapToReceipt(apiReceipt)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// our expected time (full date/time)
	expectedTime, _ := time.Parse("2006-01-02 15:04", "2022-01-02 08:13")

	if receipt.Retailer != "Walgreens" {
		t.Errorf("Expected Retailer 'Walgreens', got '%s'", receipt.Retailer)
	}
	if !receipt.PurchaseDateTime.Equal(expectedTime) {
		t.Errorf("Expected time '%v', got '%v'", expectedTime, receipt.PurchaseDateTime)
	}
	if receipt.Total != 265 {
		t.Errorf("Expected Total '265' cents, got '%d'", receipt.Total)
	}

	// our expected mapped items
	expectedItems := []rules.Item{
		{ShortDescription: "Pepsi - 12-oz", Price: 125},
		{ShortDescription: "Dasani", Price: 140},
	}

	if len(receipt.Items) != len(expectedItems) {
		t.Fatalf("Expected %d items, got %d", len(expectedItems), len(receipt.Items))
	}

	for i, expected := range expectedItems {
		actual := receipt.Items[i]
		if actual.ShortDescription != expected.ShortDescription {
			t.Errorf("Expected item %d ShortDescription '%s', got '%s'", i, expected.ShortDescription, actual.ShortDescription)
		}
		if actual.Price != expected.Price {
			t.Errorf("Expected item %d Price '%d', got '%d'", i, expected.Price, actual.Price)
		}
	}
}
