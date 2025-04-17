package main

import (
	"errors"
	"flag"
	"fmt"
	lab2 "github.com/matshp0/ArchitectureLab2"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Input file")
	outputFile      = flag.String("o", "", "Output destination")
)

func finishWithError(err error) {
	fmt.Fprintln(os.Stderr, "Error: ", err)
	os.Exit(1)
}

func main() {
	handler := lab2.ComputeHandler{}
	flag.Parse()
	if (*inputExpression != "" && *inputFile != "") ||
		(*inputExpression == "" && *inputFile == "") {
		finishWithError(errors.New("provide either -e (expression) or -f (input file), but not both"))
	}
	if *inputFile != "" {
		file, err := os.Open(*inputFile)
		handler.Reader = file
		if err != nil {
			finishWithError(err)
		}
		defer file.Close()
	}

	if *inputExpression != "" {
		handler.Reader = strings.NewReader(*inputExpression)
	}
	if *outputFile != "" {
		file, err := os.OpenFile(*outputFile, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			finishWithError(err)
		}
		handler.Writer = file
	} else {
		handler.Writer = os.Stdout
	}
	err := handler.Compute()
	if err != nil {
		finishWithError(err)
	}
}
