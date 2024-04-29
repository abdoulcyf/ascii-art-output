package cli

import (
	"errors"
	"os"
	
	"github.com/ediallocyf/output/pkg/util"
)

/*
BannerToStr converts the received string into its asccii char equivalent
*/

func BannerToStr() (string, error) {
	funcName = " --<--BannerToStr--<--"
	// get the second cliStr for reading the specified banner
	args := os.Args
	chosenBanner := args[3]
	//----------------------------------------
	patternContent, errPattern := util.ReadFileToStr(directory, chosenBanner, fileExtention)

	

	if errPattern != nil {
		funcDirectLink = "ReadFileToStr--<-- "
		funcErrorMsg = "Error : loading pattern file failed:"
		logger.Error(funcErrorMsg + funcName + funcDirectLink + errPattern.Error())
		return "", errors.New(errPattern.Error())
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
	//finalStr = strings.ReplaceAll(finalStr, "\\n", "\n")
	return finalStr, nil
}
