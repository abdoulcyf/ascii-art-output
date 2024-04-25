package cli

import (
	"errors"
	"os"

	"github.com/ediallocyf/output/pkg/util"
)

// --------------------------------------
const (
	directory          = "../assets/"
	fileExtention      = ".txt"
	chLength           = 8
	firstCh       byte = ' '
	lastCh        byte = '~'
)

//----------------------------------------

func GetFinalString() (string, error) {
	args := os.Args
	file := args[3]
	//----------------------------------------
	patternContent, errPattern := util.ReadFileToStr(directory, file, fileExtention)

	if errPattern != nil {
		errMsg = "Error : loading pattern file failed"
		return "", errors.New(errMsg)
	}
	//--------------------------------
	patternMap, errPatternMap := util.ContentToMap(patternContent, chLength)

	if errPatternMap != nil {
		return "", errors.New(errPatternMap.Error())
	}
	//--------------------------------------------
	cliStr, errClistr := util.ReadCli()

	if errClistr != nil {
		return "", errClistr
	}
	//----------------------------------------------

	finalStr := util.MapToStr(cliStr, patternMap, chLength)
	return finalStr, nil
}
