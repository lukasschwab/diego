package diego

import (
	_ "embed" // Compile-time dependency.
	"fmt"
	"log"
	"strings"
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
		return fmt.Sprintf(`lookupString(&base.%s, "%s")`, f.GoName(), f.EnvVar())
	case "int":
		return fmt.Sprintf(`%s = errors.Join(%s, lookupInt(&base.%s, "%s"))`, errName, errName, f.GoName(), f.EnvVar())
	case "bool":
		return fmt.Sprintf(`%s = errors.Join(%s, lookupBool(&base.%s, "%s"))`, errName, errName, f.GoName(), f.EnvVar())
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

// TODO: look into template composition: {{template "name" .}}

// Extract the reusable env accessors (sans package name) for inlining in
// templates.
//
//   - We could hardcode in the templates, but that would loosen the
//     relationship between env helpers and their tests.
//   - We could import in the templates, but that forces users to add a
//     dependency to their go.mod to use diego.
//
//go:embed env.go
var envAddendumRaw string
var envAddendum = strings.Split(envAddendumRaw, "TEMPLATE")[1]

//go:embed templates/base.tmpl
var baseTemplate string

// BaseTemplate for generated arg parsers.
var BaseTemplate = baseTemplate + envAddendum

//go:embed templates/struct.tmpl
var structAddendum string

// JSONTemplate is the same as the BaseTemplate, but it additionally defines the
// struct type for the parsed variables.
var JSONTemplate = BaseTemplate + structAddendum
