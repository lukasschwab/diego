package env

import (
	"fmt"
	"os"
	"strconv"
)

func LookupString(target *string, name string) {
	read, ok := os.LookupEnv(name)
	if ok {
		*target = read
	}
}

func LookupInt(target *int, name string) error {
	raw, ok := os.LookupEnv(name)
	if !ok {
		return nil
	}
	parsed, err := strconv.Atoi(raw)
	if err != nil {
		return fmt.Errorf("error parsing int environment variable '%s': %w", name, err)
	}
	*target = parsed
	return nil
}

func LookupBool(target *bool, name string) error {
	raw, ok := os.LookupEnv(name)
	if !ok {
		return nil
	}
	truthiness := raw != ""
	*target = truthiness
	return nil
}
