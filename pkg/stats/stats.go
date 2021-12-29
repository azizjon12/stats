package stats

import (
	"github.com/azizjon12/bank/pkg/types"
)

func Avg(payments []types.Payment) (amount types.Money) {
	for _, payment := range payments {
		amount += payment.Amount
	}
	return amount / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) (amount types.Money) {
	for _, payment := range payments {
		if payment.Category == category {
			amount += payment.Amount
		}
	}
	return
}
