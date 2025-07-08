package main

import (
	"log"
	"os"
)

//go:generate go run ../../cmd/diego --struct-type=ExampleVars

type ExampleVars struct {
	// --color: enable ANSI colors in CLI output
	Color bool `json:"color,omitempty"`
	// --verbose: enable verbose logging
	Verbose bool `json:"verbose,omitempty"`
	// --file: path of file to process
	File string `json:"file,omitempty"`
	// --workers: number of workers to use in parallel
	Workers int `json:"workers,omitempty"`
	// --read-only: do not write output to file
	ReadOnly bool `json:"read-only,omitempty"`
}

func main() {
	vars := new(ExampleVars)
	err := vars.Parse(os.Args[1:])

	log.Printf("Diego errors: %v", err)
	log.Printf("Parsed flags: %+v", vars)
}
