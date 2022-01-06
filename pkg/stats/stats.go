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

// CategoriesTotal returns sum of payments in each category
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}
