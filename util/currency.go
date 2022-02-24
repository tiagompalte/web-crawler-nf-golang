package util

import (
	"regexp"
	"strconv"
)

func ConvertCurrencyBrStringToFloat(str string) (float64, error) {
	re := regexp.MustCompile(`^(\d{0,3})\.?(\d{0,3})\.?(\d{0,3})\.?(\d{1,3}),(\d{2})$`)
	currencyValue := re.ReplaceAllString(str, "$1$2$3$4.$5")

	totalValue, err := strconv.ParseFloat(currencyValue, 64)
	if err != nil {
		return 0, err
	}

	return totalValue, nil
}
