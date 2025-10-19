package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// JsMinifier handles minification of JavaScript files.
type JsMinifier struct {
	*Minifier
}

// NewJsMinifier creates a new instance of JsMinifier
func NewJsMinifier(minifier *Minifier) MinifierInterface {
	return &JsMinifier{
		Minifier: minifier,
	}
}

// ReadFile reads the content of the input file
func (minifier *JsMinifier) ReadFile() error {
	content, err := os.ReadFile(minifier.InputFilename)
	if err != nil {
		return fmt.Errorf("failed to read JS file %s: %w", minifier.InputFilename, err)
	}
	minifier.Content = string(content)
	return nil
}

// Minify performs the minification process
func (minifier *JsMinifier) Minify() error {
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
func (minifier *JsMinifier) WriteFile() error {
	outputFilename := minifier.OutputFilename
	if minifier.OutputFilename == "" || outputFilename == minifier.InputFilename {
		outputFilename = strings.Replace(minifier.InputFilename, ".js", ".min.js", 1)
	}
	minifier.OutputFilename = outputFilename

	err := os.WriteFile(outputFilename, []byte(minifier.Content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write minified JS to %s: %w", outputFilename, err)
	}
	return nil
}

// jsBlockCommentRegex matches JS multi-line comments (/* ... */)
var jsBlockCommentRegex = regexp.MustCompile(`(?s)/\*.*?\*/`)

// removeComments removes all JS comments (both block and single-line)
func (minifier *JsMinifier) removeComments() error {
	// Remove block comments first
	minifier.Content = jsBlockCommentRegex.ReplaceAllString(minifier.Content, "")

	// Then remove single-line comments, preserving // inside strings
	minifier.Content = minifier.removeSingleLineComments(minifier.Content)
	return nil
}

// removeSingleLineComments removes // comments but preserves // inside strings
func (minifier *JsMinifier) removeSingleLineComments(content string) string {
	var result strings.Builder
	chars := []rune(content)
	inString := false
	stringChar := byte(0)

	for i := 0; i < len(chars); i++ {
		char := chars[i]

		// Handle string start/end
		if !inString && (char == '"' || char == '\'') {
			inString = true
			stringChar = byte(char)
			result.WriteRune(char)
			continue
		}

		if inString && byte(char) == stringChar {
			// Check if this quote is escaped
			escapeCount := 0
			for j := i - 1; j >= 0 && chars[j] == '\\'; j-- {
				escapeCount++
			}
			if escapeCount%2 == 0 {
				inString = false
				stringChar = 0
			}
			result.WriteRune(char)
			continue
		}

		// If we're in a string, keep everything
		if inString {
			result.WriteRune(char)
			continue
		}

		// Check for // comment start (not in string)
		if char == '/' && i+1 < len(chars) && chars[i+1] == '/' {
			// Skip to end of line
			for i < len(chars) && chars[i] != '\n' {
				i++
			}
			// Don't write the newline yet, let whitespace removal handle it
			if i < len(chars) {
				i--
			}
			continue
		}

		// Regular character
		result.WriteRune(char)
	}

	return result.String()
}

// Regex patterns for whitespace optimization
var jsSpaceAroundOperators = regexp.MustCompile(`\s*([{}();,=+\-*/<>!&|:\[\]])\s*`)
var jsMultipleSpaces = regexp.MustCompile(`\s+`)
var jsLineBreaks = regexp.MustCompile(`\n+`)

// removeWhiteSpace removes unnecessary spaces and line breaks in JS content
func (minifier *JsMinifier) removeWhiteSpace() error {
	str := jsSpaceAroundOperators.ReplaceAllString(minifier.Content, "$1")
	str = jsMultipleSpaces.ReplaceAllString(str, " ")
	str = jsLineBreaks.ReplaceAllString(str, "")
	str = strings.TrimSpace(str)

	minifier.Content = str
	return nil
}
