package stats

import (
	"github.com/AzizRahimov/bank/v2/pkg/types"
	"fmt"
	"testing"
	"reflect"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   53_00,
			Category: "Cat",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   51_00,
			Category: "Cat",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "Cat",
			Status:   types.StatusFail,
		},
	}

	fmt.Println(Avg(payments))

	//Output: 5200
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "auto",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   20_000_00,
			Category: "pharmacy",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "restaurant",
			Status:   types.StatusFail,
		},
	}

	inCategory := types.Category("auto")
	totalInCategory := TotalInCategory(payments, inCategory)
	fmt.Println(totalInCategory)
	//Output:  1000000
}



func TestCategoriesAvg_nill(t *testing.T) {
	var payments []types.Payment;
	

	expected := map[types.Category]types.Money{
		
	}
	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
	
	

}

func TestCategoriesAvg_empty(t *testing.T) {
	payments := []types.Payment{}

	result := CategoriesAvg(payments)

	if len(result) != 0{
		t.Error("длина не равно 0")
	}

	
}

func TestCategoriesAvg_Multiplate(t *testing.T) {
	payments := []types.Payment {
		{
			ID: 1,
			Category: "auto",
			Amount: 2000,
		},
		{
			ID: 2,
			Category: "auto",
			Amount: 2000,
		},
		{
			ID: 3,
			Category: "fun",
			Amount: 2000,
		},
		{
			ID: 4,
			Category: "mobi",
			Amount: 1000,
		},
	}
	expected := map[types.Category]types.Money{
		"auto": 2000,
		"fun": 2000,
		"mobi": 1000,
	}
	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v actual: %v", expected, result)

	}

}


func TestCategoriesAvg_FoundOne(t *testing.T) {
	payments := []types.Payment{
		{
			ID: 1,
			Category: "auto",
			Amount: 10000,
		},
	}
	expected := map[types.Category]types.Money{
		"auto": 10000,
	}
	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v got: %v", expected, result)
	}

}

func TestPeriodsDynamic_empty(t *testing.T) {
	var first = map[types.Category]types.Money{}
	var second = map[types.Category]types.Money{}
	result := PeriodsDynamic(first, second)

	if len(result) != 0 {
		t.Error("Invalid result")
	}
}
func TestPeriodsDynamic_correct(t *testing.T) {
	var first = map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	var second = map[types.Category]types.Money{
		"auto": 20,
		"food": 20,
	}

	expected := map[types.Category]types.Money{
		"auto": 10,
		"food": 0,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result:  %v", result)
	}
}

func TestPeriodsDynamic_lessSecondPeriod(t *testing.T) {
	var first = map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	var second = map[types.Category]types.Money{
		"food": 20,
	}

	expected := map[types.Category]types.Money{
		"auto": -10,
		"food": 0,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result:  %v", result)
	}
}

func TestPeriodsDynamic_lessFirstPeriod(t *testing.T) {
	var first = map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	var second = map[types.Category]types.Money{
		"auto":   10,
		"food":   25,
		"mobile": 5,
	}

	expected := map[types.Category]types.Money{
		"auto":   0,
		"food":   5,
		"mobile": 5,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result:  %v", result)
	}
	
}
