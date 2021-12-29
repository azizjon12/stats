package stats

import (
	"github.com/azizjon12/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) (amount types.Money) {
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			amount += payment.Amount
		}
	}
	return amount / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) (amount types.Money) {
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			amount += payment.Amount
		}
	}
	return
}
