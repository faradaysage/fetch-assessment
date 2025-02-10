package rules_test // prevents circular dependecy w/ mappers package, but this test makes the most sense to put here, we just happen to need to load the json w/ the mappers

import (
	mapper "fetch-assessment/mappers"
	"fetch-assessment/rules"
	"os"
	"testing"
)

func TestRulesEngineAgainstProvidedReceipts(t *testing.T) {

	tests := []struct {
		name           string
		filename       string
		expectedPoints int64
	}{
		// receipts saved from https://github.com/fetch-rewards/receipt-processor-challenge/blob/main/README.md
		{"Target Receipt", "../examples/points-test-28.json", 28},
		{"M&M Corner Market Receipt", "../examples/points-test-109.json", 109},
	}

	rulesEngine := rules.NewRulesEngine()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// load the json file
			rawJson, err := os.ReadFile(tt.filename)
			if err != nil {
				t.Fatalf("Failed to read example JSON file: %v", err)
			}

			// map the json to our business object
			receipt, err := mapper.MapToReceipt(rawJson)
			if err != nil {
				t.Fatalf("Expected no error, got: %v", err)
			}

			points := rulesEngine.CalculateTotalPoints(receipt)

			if points != tt.expectedPoints {
				t.Errorf("Expected %d points, got %d", tt.expectedPoints, points)
			}
		})
	}

}
