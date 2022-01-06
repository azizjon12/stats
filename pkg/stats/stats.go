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

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categoriesAvg, categoriesCount := make(map[types.Category]types.Money), make(map[types.Category]types.Money)

	for _, payment := range payments {
		categoriesAvg[payment.Category] += payment.Amount
		categoriesCount[payment.Category] += 1
	}

	for key, value := range categoriesAvg {
		categoriesAvg[key] = value / categoriesCount[key]
	}
	return categoriesAvg
}
