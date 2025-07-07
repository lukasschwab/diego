// Diego should generate this code from ./args.json.
package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/lukasschwab/diego/pkg/env"
)

// LukasVars generated from args.json.
type LukasVars struct {
	// --color: enable ANSI colors in CLI output
	Color bool `json:"color,omitempty"`
	// --verbose: enable verbose logging
	Verbose bool `json:"verbose,omitempty"`
	// --file: path of file to process
	File string `json:"file,omitempty"`
	// --workers: number of workers to use in parallel
	Workers int `json:"workers,omitempty"`
}

func (base *LukasVars) Parse(args []string) error {
	return errors.Join(
		base.foldEnv(),
		base.foldArgs(args),
	)
}

// TODO: probably want to use standard helpers for getting each of the three types.
func (base *LukasVars) foldEnv() error {
	var err error

	err = errors.Join(err, env.LookupBool(&base.Color, "LUKAS_COLOR"))
	err = errors.Join(err, env.LookupBool(&base.Verbose, "LUKAS_VERBOSE"))
	env.LookupString(&base.File, "LUKAS_FILE")
	err = errors.Join(err, env.LookupInt(&base.Workers, "LUKAS_WORKERS"))

	return err
}

func (base *LukasVars) foldArgs(args []string) error {
	fs := flag.NewFlagSet("LukasVars", flag.ContinueOnError)

	// NOTE: this abuses default values; the defaults should really be constant,
	// and should probably just always be the zero-values.
	//
	// In practice, these fields on the base vars might already have been set by
	// the environment.
	fs.BoolVar(&base.Color, "color", base.Color, "enable ANSI colors in CLI output")
	fs.BoolVar(&base.Verbose, "verbose", base.Verbose, "enable verbose logging")
	fs.StringVar(&base.File, "file", base.File, "path of file to process")
	fs.IntVar(&base.Workers, "workers", base.Workers, "number of workers to use in parallel")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse command line args: %w", err)
	}
	return nil
}

// TODO: add a help string helper.
