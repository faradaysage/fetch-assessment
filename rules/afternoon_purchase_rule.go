package rules

// 10 points if the time of purchase is after 2:00pm and before 4:00pm
type AfternoonPurchaseRule struct{}

func (r AfternoonPurchaseRule) CalculatePoints(receipt Receipt) int64 {
	hour, min, _ := receipt.PurchaseDateTime.Clock()
	if (hour == 14 && min >= 0) || (hour == 15 && min < 60) { // between 2:00pm (14:00 hours) and 3:59pm (15:59) [inclusively]
		return 10
	}
	return 0
}
