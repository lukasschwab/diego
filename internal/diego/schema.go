package diego

// Schema is the top-level configuration object for a set of flags.
type Schema struct {
	// Prefix for env vars; e.g. the flag --color corresponds to the env
	// variable {PREFIX}_COLOR
	EnvironmentPrefix string `json:"environmentPrefix"`
	// Flags to be parsed.
	Flags []Flag `json:"flags"`
}

// Flag describes a single command-line flag.
type Flag struct {
	// Name of the flag. This will be --{name} at the command line and
	// {PREFIX}_{NAME} in the environment.
	Name string `json:"name"`
	// Type: either bool, int, or string.
	Type string `json:"type"`
	// Description of what this flag configures.
	Description string `json:"description"`
}
