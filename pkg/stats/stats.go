package stats

import (
	"github.com/azizjon12/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) (amount types.Money) {
	numOfPmnts := 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			amount += payment.Amount
			numOfPmnts++
		}
	}
	return amount / types.Money(numOfPmnts)
}

func TotalInCategory(payments []types.Payment, category types.Category) (amount types.Money) {
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			amount += payment.Amount
		}
	}
	return
}

// FilterByCategory returns payments by the category wanted
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}
