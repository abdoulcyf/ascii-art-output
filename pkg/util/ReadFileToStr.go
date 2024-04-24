package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	//"path/filepath"
)

// ReadFileToStr reads the file content based on the given directory, banner, and file extension.
func ReadFileToStr(dir,  fileExtension string) (string, error) {
	//get chosen banner
	banner, err := PromptUserForBanner()
	if err != nil {
		return "", fmt.Errorf("failed to get banner: %v", err)
	}

	// Construct the file name based on user's choice and file extension
	fileName := filepath.Join(dir, banner+fileExtension)

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		errMsg := "->--failed to open file:--<--ReadFileToStr-<----Open-" + err.Error()
		logger.Error(errMsg)
		return "", errors.New("failed to open file")
	}
	defer file.Close() // Close the file at the end of the function

	// Read file content
	scanner := bufio.NewScanner(file)
	var fileContent string
	for scanner.Scan() {
		fileContent += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error scanning file: %v", err)
	}
	return fileContent, nil
}
