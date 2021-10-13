package stats

import "github.com/bahrom656/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	var count types.Money

	for _, payment := range payments {
		sum += payment.Amount
		count++
	}
	return sum / count
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, payment  := range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}
