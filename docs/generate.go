package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/stdlib"
)

type templateData struct {
	Title          string
	Description    string
	Example        string
	LiteralMethods map[string]object.ObjectMethod
	DefaultMethods map[string]object.ObjectMethod
}

type builtinTemplateData struct {
	Title       string
	Description string
	Example     string
	Functions   map[string]object.BuiltinFunction
	Properties  map[string]object.BuiltinProperty
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
	nil_methods := object.ListObjectMethods()[object.NIL_OBJ]
	float_methods := object.ListObjectMethods()[object.FLOAT_OBJ]
	http_methods := object.ListObjectMethods()[object.HTTP_OBJ]

	tempData := templateData{
		Title: "String",
		Example: `a = "test_string";

b = "test" + "_string";

is_true = "test" == "test";
is_false = "test" == "string";

s = "abcdef"
puts(s[2])
puts(s[-2])
puts(s[:2])
puts(s[:-2])
puts(s[2:])
puts(s[-2:])
puts(s[1:-2])

s[2] = "C"
s[-2] = "E"
puts(s)

// should output
"c"
"e"
"ab"
"abcd"
"cdef"
"ef"
"bcd"
"abCdEf"

// you can also use single quotes
'test "string" with doublequotes'

// and you can scape a double quote in a double quote string
"te\"st" == 'te"st'
`,
		LiteralMethods: string_methods,
		DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/docs/literals/string.md", tempData)

	tempData = templateData{
		Title: "Array",
		Example: `a = [1, 2, 3, 4, 5]
puts(a[2])
puts(a[-2])
puts(a[:2])
puts(a[:-2])
puts(a[2:])
puts(a[-2:])
puts(a[1:-2])

// should output
[1, 2]
[1, 2, 3]
[3, 4, 5]
[4, 5]
[2, 3]
[1, 2, 8, 9, 5]
`,
		LiteralMethods: array_methods,
		DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/docs/literals/array.md", tempData)

	tempData = templateData{
		Title: "Hash",
		Example: `people = [{"name": "Anna", "age": 24}, {"name": "Bob", "age": 99}];

// reassign of values
h = {"a": 1, 2: true}
puts(h["a"])
puts(h[2])
h["a"] = 3
h["b"] = "moo"
puts(h["a"])
puts(h["b"])
puts(h[2])h = {"a": 1, 2: true}
puts(h["a"])
puts(h[2])
h["a"] = 3
h["b"] = "moo"

// should output
1
true
3
"moo"
true`,
		LiteralMethods: hash_methods,
		DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/docs/literals/hash.md", tempData)

	tempData = templateData{
		Title:       "Boolean",
		Description: "A Boolean can represent two values: `true` and `false` and can be used in control flows.",
		Example: `true // Is the representation for truthyness
false // is it for a falsy value

a = true;
b = false;

is_true = a == a;
is_false = a == b;

is_true = a != b;`,
		LiteralMethods: boolean_methods,
		DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/docs/literals/boolean.md", tempData)

	tempData = templateData{Title: "Error", LiteralMethods: error_methods, DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/docs/literals/error.md", tempData)

	tempData = templateData{
		Title:          "File",
		Example:        `input = open("examples/aoc/2021/day-1/input").lines()`,
		LiteralMethods: file_methods,
		DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/docs/literals/file.md", tempData)

	tempData = templateData{
		Title: "Nil",
		Description: `Nil is the representation of "nothing".
	It will be returned if something returns nothing (eg. puts or an empty break/next) and can also be generated with 'nil'.`,
		LiteralMethods: nil_methods,
		DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/docs/literals/nil.md", tempData)

	tempData = templateData{
		Title: "Integer",
		Example: `a = 1;

b = a + 2;

is_true = 1 == 1;
is_false = 1 == 2;`,
		Description: `An integer can be positiv or negative and is always internally represented by a 64-Bit Integer.

To cast a negative integer a digit can be prefixed with a - eg. -456.`,
		LiteralMethods: integer_methods,
		DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/docs/literals/integer.md", tempData)

	tempData = templateData{Title: "Float", LiteralMethods: float_methods, DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/docs/literals/float.md", tempData)

	tempData = templateData{
		Title: "HTTP",
		Example: `def test()
  puts(request["body"])
  return("test")
end

HTTP.handle("/", test)

HTTP.listen(3000)

// Example request hash:
// {"protocol": "HTTP/1.1", "protocolMajor": 1, "protocolMinor": 1, "body": "servus", "method": "POST", "host": "localhost:3000", "contentLength": 6}`,
		LiteralMethods: http_methods,
		DefaultMethods: default_methods}
	create_doc("docs/templates/literal.md", "docs/docs/literals/http.md", tempData)

	// builtin module docs
	for _, module := range stdlib.Modules {
		create_doc("docs/templates/builtin.md", fmt.Sprintf("docs/docs/builtins/%s.md", module.Name), module)
	}
}

func create_doc(path string, target string, data any) bool {
	f, err := os.Create(target)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.ParseFiles(path))
	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}

	return true
}
