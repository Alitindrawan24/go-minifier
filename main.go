package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type MinifierInterface interface {
	ReadFile() error
	Minify() error
	WriteFile() error
}

type Minifier struct {
	InputFilename  string
	OutputFilename string
	Content        string
}

func main() {
	pathSrc := flag.String("src", "", "Path source file")
	pathOut := flag.String("out", "", "Path output file")
	opt := flag.String("opt", "css", "Minifier option: 'css' for CSS minification, 'js' for JavaScript minification")

	flag.Parse()
	if *pathSrc == "" {
		log.Fatal("src must be specified")
	}

	minifier := Minifier{
		InputFilename:  *pathSrc,
		OutputFilename: *pathOut,
		Content:        "",
	}

	var minifierInstance MinifierInterface

	switch *opt {
	case "css":
		minifierInstance = NewCssMinifier(&minifier)
	case "js":
		minifierInstance = NewJsMinifier(&minifier)
	}

	minifierInstance.ReadFile()
	minifierInstance.Minify()
	minifierInstance.WriteFile()

	showInformation(&minifier)
}

func showInformation(minifier *Minifier) {
	fi, err := os.Stat(minifier.InputFilename)
	if err != nil {
		panic(err)
	}

	sizeOriginal := fi.Size()
	fmt.Printf("File %s original size: %d KB\n", minifier.InputFilename, sizeOriginal/1000)

	fi, err = os.Stat(minifier.OutputFilename)
	if err != nil {
		panic(err)
	}

	sizeOutput := fi.Size()
	reduce := float64(sizeOriginal-sizeOutput) / float64(sizeOriginal) * 100
	fmt.Printf("File %s output size: %d KB (reduce by %.2f%%)\n", minifier.OutputFilename, sizeOutput/1000, reduce)
}
