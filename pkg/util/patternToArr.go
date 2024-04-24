package util

import "strings"

// ================PatternsToArr====================
func PatternsToArr(str string) ([]string, error) {
	str = strings.Trim(str, newLine)
	str = strings.ReplaceAll(str, twoLines, newLine)
	chArr := strings.Split(str, newLine)
	return chArr, nil
}
