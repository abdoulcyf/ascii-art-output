package main

import (
	"fmt"
	"strings"
)

var (
	newLine ="\n"
	twoLines = "\n\n"
)

// ================PatternsToArr====================
func PatternsToArr(str string) ([]string, error) {
	str = strings.Trim(str, newLine)
	str = strings.ReplaceAll(str, twoLines, newLine)
	chArr := strings.Split(str, newLine)
	return chArr, nil
}

func main() {
	str := "hello\n\nfirst"
	r, _ :=PatternsToArr(str)

	fmt.Println(r)
}
