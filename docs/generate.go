package main

import (
	"github.com/flipez/rocket-lang/object"
	"os"
	"text/template"
)

type templateData struct {
	Title          string
	Description    string
	Example        string
	LiteralMethods map[string]object.ObjectMethod
	DefaultMethods map[string]object.ObjectMethod
}

func main() {
	default_methods := object.ListObjectMethods()["*"]
	string_methods := object.ListObjectMethods()[object.STRING_OBJ]
	integer_methods := object.ListObjectMethods()[object.INTEGER_OBJ]
	array_methods := object.ListObjectMethods()[object.ARRAY_OBJ]
	hash_methods := object.ListObjectMethods()[object.HASH_OBJ]
	boolean_methods := object.ListObjectMethods()[object.BOOLEAN_OBJ]
	error_methods := object.ListObjectMethods()[object.ERROR_OBJ]
	file_methods := object.ListObjectMethods()[object.FILE_OBJ]
	null_methods := object.ListObjectMethods()[object.NULL_OBJ]

	tempData := templateData{Title: "String", LiteralMethods: string_methods, DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/content/docs/specification/literals/string.md", tempData)

	tempData = templateData{Title: "Array", LiteralMethods: array_methods, DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/content/docs/specification/literals/array.md", tempData)

	tempData = templateData{Title: "Hash", LiteralMethods: hash_methods, DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/content/docs/specification/literals/hash.md", tempData)

	tempData = templateData{
		Title:       "Boolean",
		Description: "A Boolean can represent two values: `true` and `false` and can be used in control flows.",
		Example: `true // Is the representation for truthyness
false // is it for a falsy value`,
		LiteralMethods: boolean_methods,
		DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/content/docs/specification/literals/boolean.md", tempData)

	tempData = templateData{Title: "Error", LiteralMethods: error_methods, DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/content/docs/specification/literals/error.md", tempData)

	tempData = templateData{Title: "File", LiteralMethods: file_methods, DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/content/docs/specification/literals/file.md", tempData)

	tempData = templateData{Title: "Null", LiteralMethods: null_methods, DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/content/docs/specification/literals/null.md", tempData)

	tempData = templateData{
		Title: "Integer",
		Description: `An integer can be positiv or negative and is always internally represented by a 64-Bit Integer.

To cast a negative integer a digit can be prefixed with a - eg. -456.`,
		LiteralMethods: integer_methods,
		DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/content/docs/specification/literals/integer.md", tempData)

}

func create_doc(path string, target string, data templateData) bool {
	paths := []string{path}

	f, err := os.Create(target)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("literal.md").ParseFiles(paths...))
	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}
	return true
}
