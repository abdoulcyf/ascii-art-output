package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// PromptUserForBanner prompts the user to choose a banner and returns the chosen banner.
func PromptUserForBanner() (string, error) {
	fmt.Print(whichBanner)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	banner := scanner.Text()

	// check validity of banner
	if banner != standard && banner != shadow && banner != thinkertoy {
		errMsg := notValidBanner
		logger.Error(errMsg)
		return "", errors.New("invalid banner")
	}

	fmt.Println("You chose:", banner)
	logger.Info(bannerChoiceInfo)

	return banner, nil
}
