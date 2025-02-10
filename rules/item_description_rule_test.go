package rules

import "testing"

// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
func TestItemDescriptionRule(t *testing.T) {
	tests := []struct {
		name     string
		desc     string
		price    int64
		expected int64
	}{
		{"Length Multiple of 3", "ABC", 100, 1},       // 20% of $1.00 = $0.20
		{"Not Multiple of 3", "AB", 100, 0},           // No points
		{"Longer Multiple of 3", "ABCDEFGHI", 501, 2}, // 20% of $5.01 = $1.002, rounded up is 2
		{"Zero price", "ABC", 0, 0},                   // 0 * 20% = 0 points
	}

	rule := ItemDescriptionRule{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt := Receipt{
				Items: []Item{{ShortDescription: tt.desc, Price: tt.price}},
			}
			points := rule.CalculatePoints(receipt)
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}
