package main

import (
	"log"
	"os"
)

//go:generate go run ../../cmd/diego ./args.json

func main() {
	vars := new(ExampleVars)
	err := vars.Parse(os.Args[1:])

	log.Printf("Diego errors: %v", err)
	log.Printf("Parsed flags: %+v", vars)
}
