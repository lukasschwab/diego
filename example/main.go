package main

import (
	"log"
	"os"
)

//go:generate go run ../cmd/diego ./args.json

func main() {
	vars := new(LUKASVars)
	err := vars.Parse(os.Args)

	log.Printf("Diego errors: %v", err)
	log.Printf("Parsed flags: %+v", vars)
}
