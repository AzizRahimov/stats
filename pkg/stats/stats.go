package stats

import (
	"github.com/AzizRahimov/bank/v2/pkg/types"
)

//Avg -
func Avg(payments []types.Payment) (money types.Money) {
	k := 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			money += payment.Amount
			k++
		}
	}
	return money / types.Money(k)
}

// TotalInCategory -
func TotalInCategory(payments []types.Payment, category types.Category) (money types.Money) {
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			money += payment.Amount
		}
	}
	return
}


func CategoriesAvg(payments[]types.Payment) map[types.Category]types.Money{
	categories := map[types.Category]types.Money{}
	var i types.Money

	for _, payment := range payments{
	categories[payment.Category] += payment.Amount

	i++
}

	for key := range categories{

	categories[key] /= i
}


	return categories
}