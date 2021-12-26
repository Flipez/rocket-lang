package main

import (
	"github.com/flipez/rocket-lang/object"
	"html/template"
	"os"
)

func main() {
	string_methods := object.ListObjectMethods()[object.STRING_OBJ]
	//for method, usage := range string_methods {
	//	fmt.Println(usage.Usage(method))
	//}

	paths := []string{"docs/templates/literal.md"}

	t := template.Must(template.New("literal.md").ParseFiles(paths...))
	err := t.Execute(os.Stdout, string_methods)
	if err != nil {
		panic(err)
	}
}
