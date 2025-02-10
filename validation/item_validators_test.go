package validation

import (
	"fetch-assessment/api"
	"testing"
)

// ensures short descriptions follow expected format
func TestShortDescriptionValidation(t *testing.T) {
	tests := []struct {
		name      string
		shortDesc string
		valid     bool
	}{
		{"Valid Description", "Mountain Dew 12PK", true},
		{"Valid With Hyphen", "Klarbrunn-12PK", true},
		{"Contains Special Characters", "Doritos!", false},
		{"Empty Description", "", false},
	}

	validator := NewShortDescriptionItemValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := api.Item{ShortDescription: tt.shortDesc}
			if validator.IsValid(item) != tt.valid {
				t.Errorf("Expected validity: %v, got: %v", tt.valid, !tt.valid)
			}
		})
	}
}

// ensures price format is valid
func TestPriceValidation(t *testing.T) {
	tests := []struct {
		name  string
		price string
		valid bool
	}{
		{"Valid Price", "6.49", true},
		{"Valid Price with Leading Zero", "0.99", true},
		{"Invalid - No Cents", "6", false},
		{"Invalid - More Than Two Decimals", "6.493", false},
		{"Invalid Characters", "6.4a", false},
		{"Empty Price", "", false},
	}

	validator := NewPriceItemValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := api.Item{Price: tt.price}
			if validator.IsValid(item) != tt.valid {
				t.Errorf("Expected validity: %v, got: %v", tt.valid, !tt.valid)
			}
		})
	}
}
