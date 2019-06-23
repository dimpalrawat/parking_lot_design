package util

import (
	"strconv"
	"strings"
)

//Function converts string to int
func StringToInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int(i)
}

//Function converts int to string
func IntToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

//Function removes spaces from string
func StripSpaces(inputString string) string {
	return strings.Replace(inputString, " ", "", -1)
}

//Function removes new lines from string
func StripNewLines(inputString string) string {
	return strings.Replace(inputString, "\n", "", -1)
}