package main

import (
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

func (minifier *CssMinifier) ReadFile() error {
	content, err := os.ReadFile(minifier.InputFilename)
	if err != nil {
		return err
	}
	minifier.Content = string(content)
	return nil
}

func (minifier *CssMinifier) Minify() error {
	err := minifier.removeComments()
	if err != nil {
		return err
	}
	err = minifier.removeWhiteSpace()
	if err != nil {
		return err
	}
	return nil
}

func (minifier *CssMinifier) WriteFile() error {
	outputFilename := minifier.OutputFilename
	if minifier.OutputFilename == "" || outputFilename == minifier.InputFilename {
		outputFilename = strings.Replace(minifier.InputFilename, ".css", ".min.css", 1)
	}
	minifier.OutputFilename = outputFilename

	err := os.WriteFile(outputFilename, []byte(minifier.Content), 0644)
	if err != nil {
		return err
	}
	return nil
}

var blockCommentRegex = regexp.MustCompile(`(?s)/\*.*?\*/`) // For single-line and multi-line comments

func (minifier *CssMinifier) removeComments() error {
	// Remove block comments
	minifier.Content = blockCommentRegex.ReplaceAllString(minifier.Content, "")
	return nil
}

var spaceAroundSymbols = regexp.MustCompile(`\s*([{}:;,])\s*`)
var spaceAroundQuotes = regexp.MustCompile(`\s+`)

func (minifier *CssMinifier) removeWhiteSpace() error {
	str := spaceAroundSymbols.ReplaceAllString(minifier.Content, "$1")
	str = spaceAroundQuotes.ReplaceAllString(str, " ")
	str = strings.ReplaceAll(str, " !important", "!important")
	str = strings.TrimSpace(str)

	minifier.Content = str
	return nil
}
