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
	first := map[types.Category]types.Money{}
	second := map[types.Category]types.Money{}

	result := PeriodsDynamic(first, second)

	if len(result)!= 0{
		t.Error("длина не равна 0")
	}

}

func TestPeriodsDynamic_addition(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 5,
	}
	second := map[types.Category]types.Money{
		"auto": 50,
		"food": 45,
		"wear": 1000,
	}
	expected := map[types.Category]types.Money{
		"auto": 40,
		"food": 40,
		"wear": 1000,
	}
	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result: expected %v got: %v", expected, result)
	}


}
func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 2000,
		"mobi": 1000,
		"food": 3000,
	}
	second := map[types.Category]types.Money{
		"auto": 3000,
		"mobi": 500,
		"food": 1000,
	}
	expected := map[types.Category]types.Money{
		"auto": 1000,
		"mobi": -500,
		"food": -2000,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result: expected %v got: %v", expected, result)
	}
	
}