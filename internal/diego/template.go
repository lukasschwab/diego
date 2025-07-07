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

// EnvLookup is exported for access from the template.
func (f TemplateFlag) EnvLookup(errName string) string {
	envVar := BuildEnvVar(f.Prefix, f.Name)
	switch f.GoType {
	case "string":
		return fmt.Sprintf(`env.LookupString(&base.%s, "%s")`, f.GoName(), envVar)
	case "int":
		return fmt.Sprintf(`%s = errors.Join(%s, env.LookupInt(&base.%s, "%s"))`, errName, errName, f.GoName(), envVar)
	case "bool":
		return fmt.Sprintf(`%s = errors.Join(%s, env.LookupBool(&base.%s, "%s"))`, errName, errName, f.GoName(), envVar)
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

//go:embed templates/diego.tmpl
var Template string
