package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ReadFileToStr reads the file content based on the given directory, fileBaseName, and file extension.
func ReadFileToStr(dir, fileBaseName,  fileExtension string) (string, error) {

	//--Construct the file name based on user's choice and file extension--
	fileName := filepath.Join(dir, fileBaseName + fileExtension)

	// Open the file
	file, errOpeningFile := os.Open(fileName)
	if errOpeningFile != nil {
		errMsg := "->--failed to open file:--<--ReadFileToStr-<----Open---<--" + errOpeningFile.Error()
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	defer file.Close() 

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
