package main

import (
	"flag"
	"io"
	lab2 "kpi-lab2"
	"os"
	"path/filepath"
	"strings"
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

	var reader io.Reader
	var writer io.Writer

	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	}

	if *inputFile != "" {
		file, _ := os.OpenFile(filepath.Join(".", *inputFile), os.O_RDONLY, 0644)
		defer file.Close()
		reader = file
	}

	if *outputFile != "" {
		writer, _ = os.Create(filepath.Join(".", *outputFile))
	} else {
		writer = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	err := handler.Compute()
	if err != nil {
		os.Exit(1)
	}
}
