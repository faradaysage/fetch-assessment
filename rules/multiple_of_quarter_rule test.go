package rules

import "testing"

// 25 points if the total is a multiple of 0.25.
func TestMultipleOfQuarterRule(t *testing.T) {
	tests := []struct {
		name     string
		total    int64
		expected int64
	}{
		{"Multiple of 0.25", 525, 25}, // $5.25 - 25 points
		{"Not a multiple", 526, 0},    // $5.26 - 0 points
		{"Exactly 0.25", 25, 25},      // $0.25 - 25 points
		{"Large multiple", 2500, 25},  // $25.00 - 25 points
	}

	rule := MultipleOfQuarterRule{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt := Receipt{Total: tt.total}
			points := rule.CalculatePoints(receipt)
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}
