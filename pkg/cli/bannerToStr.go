package cli

import (
	"errors"
	"fmt"
	//"fmt"
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
	//fmt.Printf("args[3]-----333333333---- %s\n: ", chosenBanner)
	//----------------------------------------
	patternContent, errPattern := util.ReadFileToStr(directory, chosenBanner, fileExtention)

	//fmt.Printf("patternContent----4444444444---- %s\n: ", patternContent)

	if errPattern != nil {
		funcDirectLink = "ReadFileToStr--<-- "
		funcErrorMsg = "Error : loading pattern file failed:"
		logger.Error(funcErrorMsg + funcName + funcDirectLink + errPattern.Error())
		return "", errors.New(errPattern.Error())
	}
	//--------------------------------
	patternMap, errPatternMap := util.ContentToMap(patternContent, chLength)

	//	fmt.Printf("patternMap----5555555555---- %v\n: ", patternMap)

	if errPatternMap != nil {
		return "", errors.New(errPatternMap.Error())
	}
	//--------------------------------------------
	cliStr, errClistr := util.ReadCli()
	fmt.Printf("cliStr----66666---- %v\n: ", cliStr)

	if errClistr != nil {
		return "", errClistr
	}
	//----------------------------------------------

	finalStr := util.MapToStr(cliStr, patternMap, chLength)
	//fmt.Printf("finalStr----77777---- %s\n: ", finalStr)
	return finalStr, nil
}
