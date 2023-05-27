package apis

import (
	"strconv"
	"strings"

	apis "github.com/ini8labs/admin-service/src/apis"
)

func convertDateToString(date apis.Date) string {
	dateArray := []int{date.Day, date.Month, date.Year}
	strArray := make([]string, len(dateArray))
	for i, num := range dateArray {
		strArray[i] = strconv.Itoa(num)
	}
	str := strings.Join(strArray, "/")
	return str
}

func convertWinNumbersToString(numbers []int) string {
	strArray := make([]string, len(numbers))
	for i, num := range numbers {
		strArray[i] = strconv.Itoa(num)
	}
	str := strings.Join(strArray, ",")
	return str
}
