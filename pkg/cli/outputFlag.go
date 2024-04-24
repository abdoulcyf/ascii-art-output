package cli

import (
	"errors"
	"flag"
	"os"
)

const (
	outputFileName = "banner.txt"
	outputStr      = "output"
)

func OutputFlag() (string, error) {
	// Define flag for the output filename
	outputFileN := flag.String(outputStr, outputFileName, "Output filename")
	flag.Parse()

	// Check if output filename is provided
	if *outputFileN == "" {
		return "", errors.New("please provide an output filename using the --output flag")
	}

	// Check if there are enough command-line arguments
	if flag.NArg() < 1 {
		return "", errors.New("please provide at least one string to write to the file")
	}

	// Get the content to write to the file
	content, _ := GetFinalString() // Assuming GetFinalString returns a string
	for i := 0; i < flag.NArg(); i++ {
		content += flag.Arg(i) + " "
	}

	// Create the file
	file, err := os.Create(*outputFileN)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Write content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return "", err
	}

	return *outputFileN, nil
}
