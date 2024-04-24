package util

import (
	"strings"
)

// ========================StrMaker====================================
func MapToStr(cliStr string, patternMap map[byte][]string, chLength int) string {
	var finalStr string
	cliStrArr := strings.Split(cliStr, twoLines)
	for _, strItem := range cliStrArr {
		if strItem == emptyString {
			finalStr += newLine
			continue
		}
		for row := 1; row <= chLength; row++ {
			for _, v := range strItem {
				finalStr += patternMap[byte(v)][row-1]
			}
			finalStr += newLine
		}
	}
	if ((len(cliStrArr)) == 2) && (cliStrArr[0] == emptyString) && cliStrArr[1] == emptyString {
		finalStr = finalStr[:len(finalStr)-2]
	} else {
		finalStr = finalStr[:len(finalStr)-1]
	}
	return finalStr
}
