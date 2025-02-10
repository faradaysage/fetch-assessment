package rules

import "testing"

// 5 points for every two items on the receipt.
func TestItemPairRule(t *testing.T) {
	tests := []struct {
		name     string
		items    int
		expected int64
	}{
		{"No items", 0, 0},
		{"One item", 1, 0},
		{"Two items", 2, 5},
		{"Three items", 3, 5},
		{"Four items", 4, 10},
		{"Twenty items", 20, 50},
	}

	rule := ItemPairRule{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt := Receipt{
				Items: make([]Item, tt.items),
			}
			points := rule.CalculatePoints(receipt)
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}
