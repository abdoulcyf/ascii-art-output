package util

import (
	"strings"
)

// ========================StrMaker====================================
func MapToStr(cliStr string, patternMap map[byte][]string, chLength int) string {
	var finalStr string
	cliStrArr := strings.Split(cliStr, newLine)
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
			// fmt.Println("FINALSTRING==66666666666===>", finalStr)

		}
	}
	if ((len(cliStrArr)) == 2) && (cliStrArr[0] == emptyString) && cliStrArr[1] == emptyString {
		finalStr = finalStr[:len(finalStr)-2]
	} else {
		finalStr = finalStr[:len(finalStr)-1]
	}
	//fmt.Println("FINALSTRING=====>", finalStr)
	return finalStr
}

// func MapToStr(cliStr string, patternMap map[byte][]string, chLength int) string {
//     var finalStr strings.Builder
//     cliStrArr := strings.Split(cliStr, newLine)
//     for i, strItem := range cliStrArr {
//         if i > 0 {
//             finalStr.WriteString(newLine) // Add newline between lines
//         }
//         if strItem == emptyString {
//             continue // Skip processing empty lines
//         }
//         for row := 1; row <= chLength; row++ {
//             if row > 1 {
//                 finalStr.WriteString(newLine) // Add newline between rows
//             }
//             for _, v := range strItem {
//                 finalStr.WriteString(patternMap[byte(v)][row-1])
//             }
//         }
//     }
// 	fmt.Println("FINALSTRING=====>", finalStr.String())
//     return finalStr.String()
// }
