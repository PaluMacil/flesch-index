package main

import (
	"fmt"
	"github.com/PaluMacil/flesch-index/flesch"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("No file given for analysis")
	}
	_, err := flesch.ParseFile(args[0])
	if err != nil {
		fmt.Println("cannot parse file:", err)
		os.Exit(1)
	}
}
