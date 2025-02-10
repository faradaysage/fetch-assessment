package rules

// RewardRule is the interface that all reward rules implement (strategy pattern).
type Rule interface {
	CalculatePoints(receipt Receipt) int64
}
