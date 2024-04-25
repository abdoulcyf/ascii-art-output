package cli

import (
	"flag"
	"fmt"
	"os"
)

// GenerateBannerText generates ASCII art banners from input text, allowing users to customize styling and output file name."
func GenerateBannerText() (string, error) {
	funcName = " --<--OutputFlag--<--"
	defer flag.Parse()
	//-----Define flag for the output filename-----
	outputFileN := flag.String(outputStr, outputFileName, "Output filename")

	//-----Get ascii string to write data to outputFileName-----
	asciiArt, errBannerToStr := BannerToStr()
	if errBannerToStr != nil {
		funcDirectLink = "BannerToStr--<-- "
		funcErrorMsg = "Error converting string into asscii equivalent:"
		logger.Error(funcErrorMsg + funcName + funcDirectLink + errBannerToStr.Error())
		return "", errBannerToStr
	} else {
		logger.Info("String converted into its ascii equivalence successufully")
	}

	//---write ascii string to outputFileName----------------
	if *outputFileN != "" {
		errOutputFileName := os.WriteFile(*outputFileN, []byte(asciiArt), 0644)
		if errOutputFileName != nil {
			funcDirectLink = "WriteFile--<-- "
			funcErrorMsg = "Erro writing to file:"
			logger.Error(funcErrorMsg + funcName + funcDirectLink + errOutputFileName.Error())
			return "", errOutputFileName
		} else {
			successMsg = fmt.Sprintf("Output written to %s", *outputFileN)
			logger.Info(successMsg)
		}
	}
	return *outputFileN, nil
}
