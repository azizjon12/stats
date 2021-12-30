package stats

import (
	"fmt"

	"github.com/azizjon12/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		// {
		// 	ID:       14,
		// 	Amount:   7_00,
		// 	Category: "shop",
		// 	Status:   "OK",
		// },
		// {
		// 	ID:       21,
		// 	Amount:   200_00,
		// 	Category: "restaurant",
		// 	Status:   "FAIL",
		// },
		// {
		// 	ID:       22,
		// 	Amount:   20_00,
		// 	Category: "cafe",
		// 	Status:   "OK",
		// },
		{
			ID:       14,
			Amount:   10_00,
			Category: "shop",
			Status:   "OK",
		},
		{
			ID:       21,
			Amount:   1_00,
			Category: "restaurant",
			Status:   "FAIL",
		},
		{
			ID:       22,
			Amount:   0,
			Category: "cafe",
			Status:   "OK",
		},
		{
			ID:       24,
			Amount:   3,
			Category: "cafe",
			Status:   "INPROGRESS",
		},
	}

	fmt.Println(Avg(payments))
	// Output: 250
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       314,
			Amount:   50_00,
			Category: "cafe",
			Status:   "OK",
		},
		{
			ID:       324,
			Amount:   70_00,
			Category: "restaurant",
			Status:   "OK",
		},
		{
			ID:       334,
			Amount:   100_00,
			Category: "cafe",
			Status:   "OK",
		},
		{
			ID:       334,
			Amount:   100_00,
			Category: "cafe",
			Status:   "FAIL",
		},
	}

	fmt.Println(TotalInCategory(payments, "cafe"))
	// Output: 15000
}
