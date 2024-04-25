package main

import (
	"fmt"

	"github.com/ediallocyf/output/pkg/cli"
)

func main() {
	r, _ := cli.GenerateBannerText()
	fmt.Println(r)
}
