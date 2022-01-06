package stats

import (
	"fmt"
	"reflect"
	"testing"

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
			Amount:   100_00,
			Category: "restaurant",
			Status:   "FAIL",
		},
		{
			ID:       22,
			Amount:   20_00,
			Category: "cafe",
			Status:   "OK",
		},
		{
			ID:       24,
			Amount:   30_00,
			Category: "cafe",
			Status:   "INPROGRESS",
		},
	}

	fmt.Println(Avg(payments))
	// Output: 2000
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

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "food", Amount: 100_00},
		{ID: 2, Category: "balance", Amount: 10_00},
		{ID: 3, Category: "food", Amount: 150_00},
		{ID: 4, Category: "food", Amount: 200_00},
		{ID: 5, Category: "balance", Amount: 15_00},
		{ID: 6, Category: "petrol", Amount: 200_00},
	}
	expected := map[types.Category]types.Money{
		"food":    150_00,
		"balance": 12_50,
		"petrol":  200_00,
	}

	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result expected: %v, received: %v", expected, result)
	}
}
