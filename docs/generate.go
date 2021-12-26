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
	tempData := struct {
		Methods map[string]object.ObjectMethod
	}{
		Methods: string_methods,
	}

	paths := []string{"docs/templates/literal.md"}

	t := template.Must(template.New("literal.md").ParseFiles(paths...))
	err := t.Execute(os.Stdout, tempData)
	if err != nil {
		panic(err)
	}
}
