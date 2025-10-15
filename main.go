package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// MinifierInterface defines the contract for any file minifier.
// Each minifier type (CSS, JS, etc.) must implement these methods.
type MinifierInterface interface {
	ReadFile() error
	Minify() error
	WriteFile() error
}

// Minifier holds common file data used by all specific minifiers.
type Minifier struct {
	InputFilename  string
	OutputFilename string
	Content        string
}

func main() {
	// Define command-line flags for input/output paths and minifier option.
	pathSrc := flag.String("src", "", "Path to the source file to be minified")
	pathOut := flag.String("out", "", "Path for the output minified file")
	opt := flag.String("opt", "css", "Minifier type: 'css' for CSS or 'js' for JavaScript")

	flag.Parse()

	// Validate required arguments.
	if *pathSrc == "" {
		log.Fatal("Error: --src must be specified")
	}

	// Initialize a base Minifier struct shared by all implementations.
	minifier := Minifier{
		InputFilename:  *pathSrc,
		OutputFilename: *pathOut,
		Content:        "",
	}

	// Create an instance of the specific minifier (CSS or JS).
	var minifierInstance MinifierInterface
	switch *opt {
	case "css":
		minifierInstance = NewCssMinifier(&minifier)
	case "js":
		minifierInstance = NewJsMinifier(&minifier)
	default:
		log.Fatalf("Unknown minifier option: %s. Use 'css' or 'js'.", *opt)
	}

	// Perform the minification process in three stages.
	if err := minifierInstance.ReadFile(); err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	if err := minifierInstance.Minify(); err != nil {
		log.Fatalf("Failed to minify file: %v", err)
	}

	if err := minifierInstance.WriteFile(); err != nil {
		log.Fatalf("Failed to write output: %v", err)
	}

	// Display file size reduction info after successful minification.
	showInformation(&minifier)
}

// showInformation prints file size details and reduction percentage.
func showInformation(minifier *Minifier) {
	// Get original file size.
	fi, err := os.Stat(minifier.InputFilename)
	if err != nil {
		log.Fatalf("Failed to read input file info: %v", err)
	}
	sizeOriginal := fi.Size()

	// Get output file size.
	fi, err = os.Stat(minifier.OutputFilename)
	if err != nil {
		log.Fatalf("Failed to read output file info: %v", err)
	}
	sizeOutput := fi.Size()

	// Calculate reduction percentage.
	reduce := float64(sizeOriginal-sizeOutput) / float64(sizeOriginal) * 100

	// Print statistics.
	fmt.Printf("File %s original size: %d KB\n", minifier.InputFilename, sizeOriginal/1000)
	fmt.Printf("File %s output size: %d KB (reduced by %.2f%%)\n", minifier.OutputFilename, sizeOutput/1000, reduce)
}
