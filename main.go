package main

import (
	"flag"
	"fmt"
	"github.com/PaluMacil/flesch-index/analysis"
	"github.com/PaluMacil/flesch-index/flesch"
	"os"
)

func main() {
	flagAnalysis := flag.Bool("analysis", false, "do extended analysis")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("No file given for analysis")
		os.Exit(1)
	}
	document, err := flesch.ParseFile(flag.Arg(0))
	if err != nil {
		fmt.Println("cannot parse file:", err)
		os.Exit(1)
	}

	fmt.Println("Document:", document.Name())
	fmt.Println()
	fmt.Println("Flesch Index Score:", document.Score())

	if *flagAnalysis {
		fmt.Println()
		fmt.Println("Detailed Analysis Follows:")
		report, err := analysis.Build(document)
		if err != nil {
			fmt.Println("cannot build analysis:", err)
			os.Exit(1)
		}
		fmt.Println(report.SyllableAnalysis.ChartPath)
	}
}
