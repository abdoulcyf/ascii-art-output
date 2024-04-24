package cli

import (
	"errors"
	"fmt"

	//"github.com/ediallocyf/output/pkg/cli"
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

	//----------------------------------------
	patternContent, errPattern := util.ReadFileToStr(directory, fileExtention)

	if errPattern != nil {
		errMsg = "Error : loading pattern file failed"
		return "", errors.New(errMsg)
	}
	//--------------------------------
	patternMap, errPatternMap := util.ContentToMap(patternContent, chLength)

	if errPatternMap != nil {
		//fmt.Println(errPatternMap)
		return "", errors.New(errPatternMap.Error())
	}
	//--------------------------------------------
	cliStr, errClistr := ReadCli()

	if errClistr != nil {
		//fmt.Println(errClistr)
		return "", errClistr
	}
	//----------------------------------------------

	finalStr := util.MapToStr(cliStr, patternMap, chLength)
	fmt.Println(finalStr)
	//-----------------------------------------
	return finalStr, nil
}
