package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"

	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/stdlib"
)

type templateData struct {
	Title          string
	Description    string
	Example        string
	LiteralMethods map[string]object.ObjectMethod
	DefaultMethods map[string]object.ObjectMethod
	Functions      map[string]*object.BuiltinFunction
	Properties     map[string]*object.BuiltinProperty
}

func main() {
	defaultMethods, err := loadTemplateData("object", object.ListObjectMethods()["*"])
	if err != nil {
		fmt.Printf("error loading doc for default literal methods: %s\n", err)
		return
	}

	for objType, methods := range object.ListObjectMethods() {
		if objType == "*" {
			continue
		}
		name := strings.ToLower(string(objType))

		tempData, err := loadTemplateData(name, methods)
		if err != nil {
			fmt.Printf("error loading template data for literal %s: %s\n", name, err)
			return
		}

		tempData.DefaultMethods = defaultMethods.LiteralMethods

		err = createDoc(
			"docs/templates/literal.md",
			fmt.Sprintf("docs/docs/literals/%s.md", name),
			tempData,
		)
		if err != nil {
			fmt.Printf("error creating documentation for literal %s: %s\n", name, err)
			return
		}
	}

	// builtin module docs
	for _, module := range stdlib.Modules {
		tempData, err := loadBuiltinTemplateData(module)
		if err != nil {
			fmt.Printf("error loading template data for module %s: %s\n", module.Name, err)
			return
		}
		err = createDoc(
			"docs/templates/builtin.md",
			fmt.Sprintf("docs/docs/builtins/%s.md", module.Name),
			tempData,
		)
		if err != nil {
			fmt.Printf("error creating documentation for module %s: %s\n", module.Name, err)
			return
		}
	}
}

func createDoc(path string, target string, data any) error {
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	return template.Must(template.ParseFiles(path)).Execute(f, data)
}

func loadTemplateData(name string, methods map[string]object.ObjectMethod) (*templateData, error) {
	content, err := os.ReadFile(fmt.Sprintf("docs/literals/%s.yml", name))
	if err != nil {
		return nil, err
	}

	var docData struct {
		Title       string `yaml:"title"`
		Description string `yaml:"description"`
		Example     string `yaml:"example"`
		Methods     map[string]struct {
			Description string `yaml:"description"`
			Input       string `yaml:"input"`
			Output      string `yaml:"output"`
		} `yaml:"methods"`
	}
	if err := yaml.Unmarshal(content, &docData); err != nil {
		return nil, err
	}

	tempData := templateData{
		Title:          docData.Title,
		Description:    docData.Description,
		Example:        docData.Example,
		LiteralMethods: make(map[string]object.ObjectMethod),
	}

	for name, method := range methods {
		objMethod := object.ObjectMethod{
			Layout: method.Layout,
		}
		if v, ok := docData.Methods[name]; ok {
			objMethod.Layout.Description = v.Description
			objMethod.Layout.Input = v.Input
			objMethod.Layout.Output = v.Output
		}
		tempData.LiteralMethods[name] = objMethod
	}

	return &tempData, nil
}

func loadBuiltinTemplateData(module *object.BuiltinModule) (*templateData, error) {
	content, err := os.ReadFile(fmt.Sprintf("docs/builtins/%s.yml", strings.ToLower(module.Name)))
	if err != nil {
		return nil, err
	}

	var docData struct {
		Description string `yaml:"description"`
		Example     string `yaml:"example"`
		Functions   map[string]struct {
			Description string `yaml:"description"`
			Input       string `yaml:"input"`
			Output      string `yaml:"output"`
		} `yaml:"functions"`
	}
	if err := yaml.Unmarshal(content, &docData); err != nil {
		return nil, err
	}

	tempData := templateData{
		Title:       module.Name,
		Description: docData.Description,
		Example:     docData.Example,
		Functions:   make(map[string]*object.BuiltinFunction),
		Properties:  module.Properties,
	}

	for name, function := range module.Functions {
		fn := &object.BuiltinFunction{
			Layout: function.Layout,
		}
		if v, ok := docData.Functions[name]; ok {
			fn.Layout.Description = v.Description
			fn.Layout.Input = v.Input
			fn.Layout.Output = v.Output
		}
		tempData.Functions[name] = fn
	}

	return &tempData, nil
}
