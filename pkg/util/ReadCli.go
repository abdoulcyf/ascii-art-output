package util

import (
	"errors"
	"fmt"
	"os"
	"strings"
	//"strings"
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
	fmt.Println(cliSlice)

	result := " "

	for _, strItem := range cliSlice {
		if strItem == ""{
			errMsg = "emtpy string"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
		} else {
			result += strItem
		}
	}
	// if spliStr == "" {
	// 	errMsg = "emtpy string"
	// 	logger.Error(errMsg)
	// 	return "", errors.New(errMsg)
	// }

	return strings.Join(cliSlice, "\n"), nil
}

// args := os.Args
// if len(args) != 2 {
// 	errMsg = "more than 2 input arguments"
// 	logger.Error(errMsg)
// 	return "", errors.New(errMsg)
// }

//cliStr := args[1]

// if cliStr == "" {
// 	errMsg = "emtpy string"
// 	logger.Error(errMsg)
// 	return "", errors.New(errMsg)
// }
