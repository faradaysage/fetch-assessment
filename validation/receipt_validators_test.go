package validation

import (
	"fetch-assessment/api"
	"testing"
)

// ensures retailer follows required pattern
func TestRetailerValidation(t *testing.T) {
	tests := []struct {
		name     string
		retailer string
		valid    bool
	}{
		{"Valid Retailer", "M&M Corner Market", true},
		{"Valid With Spaces", "Best Buy", true},
		{"Invalid Characters", "Shop@Home!", false},
		{"Empty Retailer", "", false},
	}

	validator := NewRetailerReceiptValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt := api.Receipt{Retailer: tt.retailer}
			if validator.IsValid(receipt) != tt.valid {
				t.Errorf("Expected validity: %v, got: %v", tt.valid, !tt.valid)
			}
		})
	}
}

// ensures valid 24-hour time format
func TestPurchaseTimeValidation(t *testing.T) {
	tests := []struct {
		name  string
		time  string
		valid bool
	}{
		{"Valid Time", "14:30", true},
		{"Invalid Format", "2:30 PM", false},
		{"Non-Time String", "Afternoon", false},
		{"Empty Time", "", false},
	}

	validator := &PurchaseTimeReceiptValidator{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt := api.Receipt{PurchaseTime: tt.time}
			if validator.IsValid(receipt) != tt.valid {
				t.Errorf("Expected validity: %v, got: %v", tt.valid, !tt.valid)
			}
		})
	}
}

// ensures receipts have at least one item
func TestReceiptItemsValidation(t *testing.T) {
	tests := []struct {
		name  string
		items []api.Item
		valid bool
	}{
		{"Valid with One Item", []api.Item{{ShortDescription: "Apple", Price: "1.00"}}, true},
		{"Valid with Multiple Items", []api.Item{{ShortDescription: "Milk", Price: "3.49"}, {ShortDescription: "Eggs", Price: "2.99"}}, true},
		{"No Items", []api.Item{}, false},
		{"One Item Valid, Other Item Invalid Short Description", []api.Item{{ShortDescription: "Milk", Price: "3.49"}, {ShortDescription: "Bad Name!", Price: "2.99"}}, false},
		{"One Item Valid, Other Item Invalid Price", []api.Item{{ShortDescription: "Milk", Price: "3.49"}, {ShortDescription: "Eggs", Price: "$100"}}, false},
	}

	validator := &ItemsReceiptValidator{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt := api.Receipt{Items: tt.items}
			if validator.IsValid(receipt) != tt.valid {
				t.Errorf("Expected validity: %v, got: %v", tt.valid, !tt.valid)
			}
		})
	}
}

// ensures valid receipt total
func TestTotalValidation(t *testing.T) {
	tests := []struct {
		name  string
		time  string
		valid bool
	}{
		{"Valid Price", "6.49", true},
		{"Valid Price with Leading Zero", "0.99", true},
		{"Invalid - No Cents", "6", false},
		{"Invalid - More Than Two Decimals", "6.493", false},
		{"Invalid Characters", "6.4a", false},
		{"Empty Price", "", false},
	}

	validator := NewTotalReceiptValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt := api.Receipt{Total: tt.time}
			if validator.IsValid(receipt) != tt.valid {
				t.Errorf("Expected validity: %v, got: %v", tt.valid, !tt.valid)
			}
		})
	}
}
