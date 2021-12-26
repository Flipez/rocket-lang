package main

import (
	"github.com/flipez/rocket-lang/object"
	"os"
	"text/template"
)

func main() {
	string_methods := object.ListObjectMethods()[object.STRING_OBJ]
	//for method, usage := range string_methods {
	//	fmt.Println(usage.Usage(method))
	//}
	tempData := struct {
		Methods map[string]object.ObjectMethod
	}{
		Methods: string_methods,
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
