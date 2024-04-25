package cli

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
)



func OutputFlag() (string, error) {
	// Define flag for the output filename
	outputFileN := flag.String(outputStr, outputFileName, "Output filename")
	flag.Parse()

	
	asciiArt, _ := GetFinalString()

	if *outputFileN != "" {
		err := ioutil.WriteFile(*outputFileN, []byte(asciiArt), 0644)
		if err != nil {
			errMsg = "Error writing to file:" + err.Error()
			return "", errors.New(errMsg)
		}
		fmt.Printf("Output written to %s\n", *outputFileN)
	} else {
		// If no output file specified, print ASCII art to stdout
		fmt.Println(asciiArt)
	}
	return *outputFileN, nil
}

func WriteFile(s string, b []byte, i int) {
	panic("unimplemented")
}
