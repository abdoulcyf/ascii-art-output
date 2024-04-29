package cli

import (
	"flag"
	"fmt"
	"os"
)

// GenerateBannerText generates ASCII art banners from input text, allowing users to customize styling and output file name."
func GenerateBannerText(defaultOutputFileName string) (string, error) {
	funcName := "GenerateBannerText"
	var outputFileName string

	// Define flag for the output filename
	flag.StringVar(&outputFileName, "output", defaultOutputFileName, "Output filename")

	// Parse flags
	flag.Parse()

	// Replace "\\n" with actual newline character

	// Get ASCII string to write data to outputFileName
	asciiArt, errBannerToStr := BannerToStr()
	if errBannerToStr != nil {
		funcDirectLink := "BannerToStr"
		funcErrorMsg := "Error converting string into ASCII equivalent: "
		logger.Error(funcErrorMsg + funcName + funcDirectLink + errBannerToStr.Error())
		return "", errBannerToStr
	}
	logger.Info("String converted into its ASCII equivalent successfully")

	// Write ASCII string to outputFileName
	if outputFileName != "" {
		errOutputFileName := os.WriteFile(outputFileName, []byte(asciiArt), 0644)
		if errOutputFileName != nil {
			funcDirectLink := "WriteFile"
			funcErrorMsg := "Error writing to file: "
			logger.Error(funcErrorMsg + funcName + funcDirectLink + errOutputFileName.Error())
			return "", errOutputFileName
		}
		successMsg := fmt.Sprintf("Output written to %s", outputFileName)
		logger.Info(successMsg)
	}
	return outputFileName, nil
}
