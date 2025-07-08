package diego

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BODGE: the comment on the following line separates the helpers to be inlined
// in template.go from the go metadata that makes this code testable in the
// current package: the package and import declarations above.
//
// TEMPLATE

// lookupString in environment; write to target if it's set.
func lookupString(target *string, name string) {
	read, ok := os.LookupEnv(name)
	if ok {
		*target = read
	}
}

// lookupInt in environment; write to target if it's set and parseable as a
// decimal int.
func lookupInt(target *int, name string) error {
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

// lookupBool in environment; write to target if it's set.
func lookupBool(target *bool, name string) error {
	raw, ok := os.LookupEnv(name)
	if !ok {
		return nil
	}
	truthiness := raw != "" && strings.ToLower(raw) != "false"
	*target = truthiness
	return nil
}
