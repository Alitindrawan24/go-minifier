package main

import (
	"flag"
	"log"
)

type MinifierInterface interface {
	ReadFile()
	RemoveWhiteSpace()
	RemoveComments()
	WriteFile()
}

type Minifier struct {
	InputFilename  string
	OutputFilename string
	Content        string
}

func main() {
	pathSrc := flag.String("src", "", "Path source file")
	pathOut := flag.String("out", "", "Path output file")

	flag.Parse()
	if *pathSrc == "" {
		log.Fatal("src must be specified")
	}

	minifier := Minifier{
		InputFilename:  *pathSrc,
		OutputFilename: *pathOut,
		Content:        "",
	}

	cssMinifier := NewCssMinifier(minifier)
	cssMinifier.ReadFile()
	cssMinifier.RemoveWhiteSpace()
	cssMinifier.RemoveComments()
	cssMinifier.WriteFile()
}
