package rules

import (
	"testing"
)

// 50 points if the total is a round dollar amount with no cents.
func TestRoundDollarRule(t *testing.T) {
	tests := []struct {
		name     string
		total    int64
		expected int64
	}{
		{"Round Dollar Amount", 500, 50}, // $5.00 - should get 50 points
		{"Not a Round Dollar", 525, 0},   // $5.25 - should get 0 points
		{"Exact Round Dollar", 1000, 50}, // $10.00 - should get 50 points
	}

	rule := RoundDollarRule{}

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
