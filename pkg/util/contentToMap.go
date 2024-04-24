package util

import "errors"

// =======================ContentToMap===================================
func ContentToMap(content string, chLength int) (map[byte][]string, error) {
	contentArr, errPatternsToArr := PatternsToArr(content)
	if errPatternsToArr != nil {
		errMsg = "--ContentToMap-->--PatternsToArr----->" + errPatternsToArr.Error()
		logMsg = "Error in converting partern to an array"
		logger.Error(logMsg + errMsg)
		return nil, errors.New(errMsg)
	}
	contentMap, err := ToMap(contentArr, chLength)
	return contentMap, err
}