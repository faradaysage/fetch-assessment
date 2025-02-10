package mapper

import (
	"fetch-assessment/rules"
	"os"
	"testing"
	"time"
)

// test that we are properly mapping the json in example/morning-receipt.json to our business object
func TestJsonReceiptMapperForMorningReceipt(t *testing.T) {
	// load the json file
	rawJson, err := os.ReadFile("../examples/morning-receipt.json")
	if err != nil {
		t.Fatalf("Failed to read example JSON file: %v", err)
	}

	// map the json to our business object
	receipt, err := MapToReceipt(rawJson)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// "purchaseDate": "2022-01-02",
	// "purchaseTime": "08:13",
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

	// {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
	// {"shortDescription": "Dasani", "price": "1.40"}
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

// test that we are properly mapping the json in example/simple-receipt.json to our business object
func TestJsonReceiptMapperForSimpleReceipt(t *testing.T) {
	// load the json file
	rawJson, err := os.ReadFile("../examples/simple-receipt.json")
	if err != nil {
		t.Fatalf("Failed to read example JSON file: %v", err)
	}

	// create the mapper
	mapper := JsonReceiptMapper{
		ItemMapper: &JsonItemMapper{},
	}

	// map the json to our business object
	receipt, err := mapper.ToDomain(rawJson)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// "purchaseDate": "2022-01-02",
	// "purchaseTime": "13:13",
	expectedTime, _ := time.Parse("2006-01-02 15:04", "2022-01-02 13:13")

	if receipt.Retailer != "Target" {
		t.Errorf("Expected Retailer 'Target', got '%s'", receipt.Retailer)
	}
	if !receipt.PurchaseDateTime.Equal(expectedTime) {
		t.Errorf("Expected time '%v', got '%v'", expectedTime, receipt.PurchaseDateTime)
	}
	if receipt.Total != 125 {
		t.Errorf("Expected Total '125' cents, got '%d'", receipt.Total)
	}

	// {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
	expectedItems := []rules.Item{
		{ShortDescription: "Pepsi - 12-oz", Price: 125},
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
