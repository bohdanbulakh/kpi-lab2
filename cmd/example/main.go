package main

import (
	"flag"
	"os"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Input file with expression to compute")
	outputFile      = flag.String("o", "", "File with result")
)

func main() {
	flag.Parse()

	if *inputExpression != "" && *inputFile != "" {
		_, _ = os.Stderr.WriteString("You cannot use both -e and -f\n")
		os.Exit(1)
	}

	if *inputExpression == "" && *inputFile == "" {
		_, _ = os.Stderr.WriteString("You must use one of these arguments: -e, -f\n")
		os.Exit(1)
	}
}
