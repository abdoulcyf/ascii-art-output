package util

import "errors"

// ========================ToMap======================================
func ToMap(strArr []string, chLength int) (map[byte][]string, error) {

	// check if there are enough lines
	if len(strArr) < int(linesNumber)*chLength {
		errMsg = "not enough lines:---ToMap-<----linesNumber------"
		logger.Error(errMsg)
		return nil, errors.New("not enough lines")
	}

	//
	resultMap := make(map[byte][]string)
	for i := firstChar; i <= lastChar; i++ {
		startIndex := int(i-firstChar) * (chLength)
		endIndex := startIndex + chLength
		resultMap[byte(i)] = strArr[startIndex:endIndex]
	}
	//fmt.Println("===Map====", resultMap)
	return resultMap, nil
}
