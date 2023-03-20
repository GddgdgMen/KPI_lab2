package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	lab2 "github.com/GddgdgMen/KPI_lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inFile          = flag.String("f", "", "Input file name")
	outFile         = flag.String("o", "", "Output file name")
)

func main() {
	flag.Parse()

	if *inputExpression == "" && *inFile == "" {
		log.Fatal("no expression provided. use -e \"{expression}\" or -f {file with expression}")
	}

	var reader io.Reader

	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else {
		file, err := os.Open(*inFile)
		if err != nil {
			log.Fatal("no such file")
		}
		reader = file
		defer file.Close()
	}

	var writer io.Writer

	if *outFile != "" {
		file, err := os.Create(*outFile)
		if err != nil {
			log.Fatal("something went wrong while creating file")
		}

		writer = file
		defer file.Close()
	} else {
		writer = &Writer{}
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	err := handler.Compute()

	if err != nil {
		fmt.Print(lab2.ThrowError())
	}

}

type Writer struct{}

func (w *Writer) Write(data []byte) (n int, err error) {
	fmt.Println(string(data))
	return len(data), nil
}
