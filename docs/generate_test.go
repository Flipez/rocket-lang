package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/stdlib"
)

// docKeys reads the method or function names documented in a YAML file. It
// returns nil when the file is missing, which the caller reports as every
// name being undocumented.
func docKeys(t *testing.T, path, section string) map[string]bool {
	t.Helper()

	content, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("%s: %s", path, err)
		return nil
	}

	var doc struct {
		Methods   map[string]any `yaml:"methods"`
		Functions map[string]any `yaml:"functions"`
	}
	if err := yaml.Unmarshal(content, &doc); err != nil {
		t.Fatalf("%s: %s", path, err)
	}

	source := doc.Methods
	if section == "functions" {
		source = doc.Functions
	}

	keys := make(map[string]bool, len(source))
	for name := range source {
		keys[name] = true
	}

	return keys
}

// reportUndocumented fails with every missing name at once, sorted, rather
// than stopping at the first -- during a rename the whole list is what you
// want to see.
func reportUndocumented(t *testing.T, path string, have map[string]bool, want []string) {
	t.Helper()

	var missing []string
	for _, name := range want {
		if !have[name] {
			missing = append(missing, name)
		}
	}
	sort.Strings(missing)

	if len(missing) > 0 {
		t.Errorf("%s is missing %d entries: %s", path, len(missing), strings.Join(missing, " "))
	}
}

// TestEveryMethodIsDocumented guards the rename. generate.go skips a method it
// finds no documentation for, without complaining, so a renamed method quietly
// loses its description on the website instead of failing the build.
func TestEveryMethodIsDocumented(t *testing.T) {
	for objectType, methods := range object.ListObjectMethods() {
		name := "object"
		if objectType != "*" {
			name = strings.ToLower(string(objectType))
		}

		path := fmt.Sprintf("literals/%s.yml", name)

		names := make([]string, 0, len(methods))
		for method := range methods {
			names = append(names, method)
		}

		reportUndocumented(t, path, docKeys(t, path, "methods"), names)
	}
}

// TestEveryBuiltinFunctionIsDocumented is the same guard for the stdlib modules.
func TestEveryBuiltinFunctionIsDocumented(t *testing.T) {
	for _, module := range stdlib.Modules {
		path := fmt.Sprintf("builtins/%s.yml", strings.ToLower(module.Name))

		names := make([]string, 0, len(module.Functions))
		for function := range module.Functions {
			names = append(names, function)
		}

		reportUndocumented(t, path, docKeys(t, path, "functions"), names)
	}
}
