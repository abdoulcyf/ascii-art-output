package main

import (
	"fmt"

	"github.com/ediallocyf/output/pkg/cli"
)

const defaultOutputFileName = "banner.txt"

func main() {
	bannerText, err := cli.GenerateBannerText(defaultOutputFileName)
	if err != nil {
		errMsg = "--main-->--GenerateBannerText----->" + err.Error()
		logMsg = "Error Generating banner text: "
		logger.Error(logMsg + errMsg)
		return
	} else {
		logger.Info("Banner Text generated successfully")
	}
	fmt.Println(bannerText)
}
