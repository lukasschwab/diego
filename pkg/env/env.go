package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LookupString in environment; write to target if it's set.
func LookupString(target *string, name string) {
	read, ok := os.LookupEnv(name)
	if ok {
		*target = read
	}
}

// LookupInt in environment; write to target if it's set and parseable as a
// decimal int.
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

// LookupBool in environment; write to target if it's set.
func LookupBool(target *bool, name string) error {
	raw, ok := os.LookupEnv(name)
	if !ok {
		return nil
	}
	truthiness := raw != "" && strings.ToLower(raw) != "false"
	*target = truthiness
	return nil
}
