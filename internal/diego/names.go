package diego

import (
	"fmt"
	"regexp"
	"strings"
)

const validPrefixRegex = `^[\dA-z_]*$`

type Prefix string

// ValidatePrefix: should be alphanumeric and `_`-delimited.
func ValidatePrefix(schemaPrefix string) (Prefix, error) {
	if match, err := regexp.Match(validPrefixRegex, []byte(schemaPrefix)); err != nil {
		return "", err
	} else if !match {
		return "", fmt.Errorf("invalid environment prefix '%s'; should match %s", schemaPrefix, validPrefixRegex)
	}
	return Prefix(strings.ToUpper(schemaPrefix)), nil
}

// BuildEnvVar from prefix and name. If you have prefix "FOO" and name `bar-baz`
// you get the env var `FOO_BAR_BAZ`.
func BuildEnvVar(prefix Prefix, name string) string {
	return fmt.Sprintf("%s_%s", prefix, strings.ToUpper(strings.ReplaceAll(name, "-", "_")))
}

// BuildGoName from config name. If you have the flag `bar-baz`, the Go name is
// `BarBaz`.
func BuildGoName(name string) string {
	parts := strings.Split(name, "-")
	for i := range parts {
		parts[i] = strings.Title(strings.ToLower(parts[i]))
	}
	return strings.Join(parts, "")
}
