package cli 

import (
	"errors"
	"os"
)

/*
write a function that read the arguments of command-lines and return
- it must have just 1 arguments (filename)
- err : more than 1 arguments
*/
func ReadCli() (string, error) {
	args := os.Args
	if len(args) != 2 {
		errMsg = "more than 1 input arguments"
		 logger.Error(errMsg)
		return "", errors.New(errMsg)
	}


	cliStr := args[1]

	if cliStr=="" {
		errMsg = "emtpy string"
		logger.Error(errMsg)
		return "",errors.New(errMsg)
	}

	return cliStr, nil
}
