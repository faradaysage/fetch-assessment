package rules

import (
	"testing"
	"time"
)

// 10 points if the time of purchase is after 2:00pm and before 4:00pm
func TestAfternoonPurchaseRule(t *testing.T) {
	tests := []struct {
		name     string
		timeStr  string
		expected int64
	}{
		{"Before 2PM", "2025-02-09 13:59", 0},       // 1:59 PM - No points
		{"Exactly 2PM", "2025-02-09 14:00", 10},     // 2:00 PM - 10 points
		{"Between 2PM-4PM", "2025-02-09 15:00", 10}, // 3:00 PM - 10 points
		{"Exactly 4PM", "2025-02-09 16:00", 0},      // 4:00 PM - No points
	}

	rule := AfternoonPurchaseRule{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dateTime, _ := time.Parse("2006-01-02 15:04", tt.timeStr)
			receipt := Receipt{PurchaseDateTime: dateTime}
			points := rule.CalculatePoints(receipt)
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}
