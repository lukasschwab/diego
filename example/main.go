package main

import (
	"log"
	"os"
)

func main() {
	vars := new(LukasVars)
	err := vars.Parse(os.Args)

	log.Printf("Diego errors: %v", err)
	log.Printf("Parsed flags: %+v", vars)
}
