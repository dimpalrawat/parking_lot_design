package util

import "strconv"

//Function converts string to int
func StringToInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int(i)
}

//Function converts int to string
func IntToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}