package diego

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

// FromAst builds a TemplateSchema from a Go source file rather than a JSON
// source: structName is the name of the struct defining the argument shape.
// For example, you might provide struct name `ExampleVars` to generate parse
// logic for this struct:
//
//	// ExampleVars is my command line variables struct.
//	type ExampleVars struct {
//		// --color: enable ANSI colors in CLI output
//		Color bool `json:"color,omitempty"`
//		// --verbose: enable verbose logging
//		Verbose bool `json:"verbose,omitempty"`
//		// --file: path of file to process
//		File string `json:"file,omitempty"`
//		// --workers: number of workers to use in parallel
//		Workers int `json:"workers,omitempty"`
//		// --read-only: do not write output to file
//		ReadOnly bool `json:"read-only,omitempty"`
//	}
func FromAst(goFilename, structName string) (*TemplateSchema, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, goFilename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse source file: %w", err)
	}

	var ts *TemplateSchema
	ast.Inspect(node, func(n ast.Node) bool {
		if ts != nil {
			return false // Stop searching once we've found our struct.
		}
		// We're looking for a type declaration.
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		// With the correct name.
		if typeSpec.Name.Name != structName {
			return true
		}

		// That is a struct.
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		ts = &TemplateSchema{
			Package:    node.Name.Name,
			StructName: structName,
			Source:     goFilename,
			Flags:      make([]TemplateFlag, len(structType.Fields.List)),
		}

		for i, field := range structType.Fields.List {
			flagName := getFlagName(field)
			ts.Flags[i] = TemplateFlag{
				Name:        flagName,
				Description: parseDescription(field.Doc.Text()),
				GoType:      fmt.Sprintf("%s", field.Type),
			}
		}

		prefix, err := ValidatePrefix(strings.TrimSuffix(structName, "Vars"))
		if err != nil {
			// This should be a validation error, but for now we'll just log it.
			// TODO: Return an error.
			fmt.Printf("Error validating prefix: %v\n", err)
		}
		ts.Prefix = prefix
		for i := range ts.Flags {
			ts.Flags[i].Prefix = prefix
		}

		return false
	})

	if ts == nil {
		return nil, fmt.Errorf("struct '%s' not found in %s", structName, goFilename)
	}
	return ts, nil
}

func getFlagName(field *ast.Field) string {
	if field.Tag == nil {
		return ""
	}
	tag := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])
	jsonTag := tag.Get("json")
	return strings.Split(jsonTag, ",")[0]
}

// parseDescription from a colon-separated comment like the following:
//
//	// --color: enable ANSI colors in CLI output
func parseDescription(doc string) string {
	// The doc string includes the comment markers. Trim them.
	doc = strings.TrimSpace(doc)
	// The description is after the flag name and colon.
	_, after, found := strings.Cut(doc, ":")
	if !found {
		return ""
	}
	return strings.TrimSpace(after)
}
