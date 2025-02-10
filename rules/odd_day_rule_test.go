package rules

import (
	"testing"
	"time"
)

// 6 points if the day in the purchase date is odd
func TestOddDayRule(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		expected int64
	}{
		{"Odd Day", "2025-02-09", 6},  // February 9th - 6 points
		{"Even Day", "2025-02-10", 0}, // February 10th - 0 points
	}

	rule := OddDayRule{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date, _ := time.Parse("2006-01-02", tt.date)
			receipt := Receipt{PurchaseDateTime: date}
			points := rule.CalculatePoints(receipt)
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}
