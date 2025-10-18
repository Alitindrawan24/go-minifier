package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// CssMinifier handles minification of CSS files.
type CssMinifier struct {
	*Minifier
}

// NewCssMinifier creates a new instance of CssMinifier
func NewCssMinifier(minifier *Minifier) MinifierInterface {
	return &CssMinifier{
		Minifier: minifier,
	}
}

// ReadFile reads the content of the input file
func (minifier *CssMinifier) ReadFile() error {
	content, err := os.ReadFile(minifier.InputFilename)
	if err != nil {
		return fmt.Errorf("failed to read CSS file %s: %w", minifier.InputFilename, err)
	}
	minifier.Content = string(content)
	return nil
}

// Minify performs the minification process
func (minifier *CssMinifier) Minify() error {
	// Step 1: Remove all comments
	err := minifier.removeComments()
	if err != nil {
		return err
	}

	// Step 2: Remove unnecessary whitespace
	err = minifier.removeWhiteSpace()
	if err != nil {
		return err
	}

	return nil
}

// WriteFile writes the minified content to the output file
func (minifier *CssMinifier) WriteFile() error {
	outputFilename := minifier.OutputFilename
	if minifier.OutputFilename == "" || outputFilename == minifier.InputFilename {
		outputFilename = strings.Replace(minifier.InputFilename, ".css", ".min.css", 1)
	}
	minifier.OutputFilename = outputFilename

	err := os.WriteFile(outputFilename, []byte(minifier.Content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write minified CSS to %s: %w", outputFilename, err)
	}
	return nil
}

// blockCommentRegex matches CSS block comments (/* ... */)
var blockCommentRegex = regexp.MustCompile(`(?s)/\*.*?\*/`)

// removeComments removes all CSS block comments (/* ... */) from the content.
func (minifier *CssMinifier) removeComments() error {
	minifier.Content = blockCommentRegex.ReplaceAllString(minifier.Content, "")
	return nil
}

// Regex patterns for whitespace optimization
var spaceAroundSymbols = regexp.MustCompile(`\s*([{}:;,])\s*`)
var spaceAroundQuotes = regexp.MustCompile(`\s+`)

// removeWhiteSpace removes unnecessary spaces around symbols and inside the content.
func (minifier *CssMinifier) removeWhiteSpace() error {
	str := spaceAroundSymbols.ReplaceAllString(minifier.Content, "$1")
	str = spaceAroundQuotes.ReplaceAllString(str, " ")
	str = strings.ReplaceAll(str, " !important", "!important")
	str = strings.TrimSpace(str)

	minifier.Content = str
	return nil
}
