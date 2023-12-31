package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)

type CssMinifier struct {
	*Minifier
}

func NewCssMinifier(minifier *Minifier) MinifierInterface {
	return &CssMinifier{
		Minifier: minifier,
	}
}

func (minifier *CssMinifier) ReadFile() {
	content, err := os.ReadFile(minifier.InputFilename)
	if err != nil {
		log.Fatal(err.Error())
	}
	minifier.Content = string(content)
}

func (minifier *CssMinifier) Minify() {
	minifier.removeComments()
	minifier.removeWhiteSpace()
}

func (minifier *CssMinifier) WriteFile() {
	outputFilename := minifier.OutputFilename
	if minifier.OutputFilename == "" {
		outputFilename = strings.Replace(minifier.InputFilename, ".css", ".min.css", 1)
	}
	minifier.OutputFilename = outputFilename

	err := os.WriteFile(outputFilename, []byte(minifier.Content), 0644)
	if err != nil {
		panic(err)
	}
}

func (minifier *CssMinifier) removeComments() {
	// Remove block comments
	blockCommentRegex := regexp.MustCompile(`/\*.*?\*/`)
	minifier.Content = blockCommentRegex.ReplaceAllString(minifier.Content, "")
}

func (minifier *CssMinifier) removeWhiteSpace() {
	// Remove extra white spaces and newlines
	str := strings.Join(strings.Fields(minifier.Content), " ")
	str = strings.ReplaceAll(str, "{ ", "{")
	str = strings.ReplaceAll(str, " {", "{")
	str = strings.ReplaceAll(str, "} ", "}")
	str = strings.ReplaceAll(str, "; ", ";")
	str = strings.ReplaceAll(str, ": ", ":")
	str = strings.ReplaceAll(str, ", ", ",")
	str = strings.ReplaceAll(str, ";}", "}")

	minifier.Content = str
}
