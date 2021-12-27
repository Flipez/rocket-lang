package main

import (
	"github.com/flipez/rocket-lang/object"
	"os"
	"text/template"
)

func main() {
	default_methods := object.ListObjectMethods()["*"]
	string_methods := object.ListObjectMethods()[object.STRING_OBJ]

	tempData := struct {
		StringMethods  map[string]object.ObjectMethod
		DefaultMethods map[string]object.ObjectMethod
	}{
		StringMethods:  string_methods,
		DefaultMethods: default_methods,
	}

	paths := []string{"docs/templates/literal.md"}

	f, err := os.Create("docs/content/docs/specification/literals/string.md")
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("literal.md").ParseFiles(paths...))
	err = t.Execute(f, tempData)
	if err != nil {
		panic(err)
	}
}
