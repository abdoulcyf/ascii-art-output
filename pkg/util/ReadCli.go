package util

import (
	"errors"
	"os"
	"strings"
)

/*
write a function that read the arguments of command-lines and return
- it must have just 2 arguments
- err : more than 1 arguments
*/
func ReadCli() (string, error) {
	args := os.Args
	if len(args) < 3 {
		errMsg = "more than 2 input arguments"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	cliStr := args[2]

	cliSlice := strings.Split(cliStr, "\\n")
	var result []string 

	for _, strItem := range cliSlice{
		strItem = strings.Trim(strItem, "\n")
		if strItem == ""{
			errMsg = "Empty string: ---ReadCli--<<---strItem--"
			logger.Error(errMsg)
			return "", errors.New(errMsg)
		} else {
			result = append(result, strItem)
		}
	}
	return strings.Join(result, "\n"), nil
}

