package diego

import (
	_ "embed" // Compile-time dependency.
	"fmt"
	"log"
)

// TemplateSchema is an intermediate derived from Schema for use with Template.
type TemplateSchema struct {
	Package    string
	StructName string
	Source     string
	Flags      []TemplateFlag
	Prefix     Prefix
}

// TemplateFlag is an intermediate derived from Flag for use with Template.
type TemplateFlag struct {
	Name        string
	Description string
	GoType      string
	Prefix      Prefix
}

// GoName is exported for access from the template.
func (f TemplateFlag) GoName() string {
	return BuildGoName(f.Name)
}

func (f TemplateFlag) EnvVar() string {
	return BuildEnvVar(f.Prefix, f.Name)
}

// EnvLookup is exported for access from the template.
func (f TemplateFlag) EnvLookup(errName string) string {
	switch f.GoType {
	case "string":
		return fmt.Sprintf(`env.LookupString(&base.%s, "%s")`, f.GoName(), f.EnvVar())
	case "int":
		return fmt.Sprintf(`%s = errors.Join(%s, env.LookupInt(&base.%s, "%s"))`, errName, errName, f.GoName(), f.EnvVar())
	case "bool":
		return fmt.Sprintf(`%s = errors.Join(%s, env.LookupBool(&base.%s, "%s"))`, errName, errName, f.GoName(), f.EnvVar())
	default:
		log.Fatalf("Unsupported go type '%v'", f.GoType)
		return ""
	}
}

// FlagVar is exported for access from the template.
func (f TemplateFlag) FlagVar() string {
	switch f.GoType {
	case "string":
		return "StringVar"
	case "int":
		return "IntVar"
	case "bool":
		return "BoolVar"
	default:
		log.Fatalf("Unsupported go type '%v'", f.GoType)
		return ""
	}
}

//go:embed templates/base.tmpl
var BaseTemplate string

//go:embed templates/struct.tmpl
var structAddendum string

// JSONTemplate is the same as the BaseTemplate, but it additionally defines the
// struct type for the parsed variables.
var JSONTemplate = BaseTemplate + structAddendum
