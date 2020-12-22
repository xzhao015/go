package ch2

import (
	"regexp"
	"strings"
)

//ReformatDate is a function.
func ReformatDate(date string) string {
	monthMap := map[string]string{
		"Jan": "1",
		"Feb": "2",
		"Mar": "3",
		"Apr": "4",
		"May": "5",
		"Jun": "6",
		"Jul": "7",
		"Aug": "8",
		"Sep": "9",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12"}
	arr := strings.Fields(date)
	var day, month, year = arr[0], arr[1], arr[2]
	dayReg := regexp.MustCompile("[0-9]{1,2}")
	dayString := dayReg.Find([]byte(day))
	outputList := []string{year, monthMap[month], string(dayString)}
	output := strings.Join(outputList, "-")
	return output
}
