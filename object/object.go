package object

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/flipez/rocket-lang/ast"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
	InvokeMethod(method string, env Environment, args ...Object) Object
}

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
	FILE_OBJ         = "FILE"
)

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
func (h *Hash) InvokeMethod(method string, env Environment, args ...Object) Object { return nil }

type Hashable interface {
	HashKey() HashKey
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType                                                   { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string                                                    { return "builtin function" }
func (b *Builtin) InvokeMethod(method string, env Environment, args ...Object) Object { return nil }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
func (f *Function) InvokeMethod(method string, env Environment, args ...Object) Object { return nil }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType                                                   { return ERROR_OBJ }
func (e *Error) Inspect() string                                                    { return "ERROR: " + e.Message }
func (e *Error) InvokeMethod(method string, env Environment, args ...Object) Object { return nil }

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string                                                    { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType                                                   { return INTEGER_OBJ }
func (i *Integer) InvokeMethod(method string, env Environment, args ...Object) Object { return nil }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) InvokeMethod(method string, env Environment, args ...Object) Object {
	return nil
}

type Null struct{}

func (n *Null) Type() ObjectType                                                   { return NULL_OBJ }
func (n *Null) Inspect() string                                                    { return "null" }
func (n *Null) InvokeMethod(method string, env Environment, args ...Object) Object { return nil }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
func (rv *ReturnValue) InvokeMethod(method string, env Environment, args ...Object) Object {
	return nil
}

type File struct {
	Filename string
	Reader   *bufio.Reader
	Writer   *bufio.Writer
	Handle   *os.File
}

func (f *File) Type() ObjectType { return FILE_OBJ }
func (f *File) Inspect() string  { return fmt.Sprintf("<file:%s>", f.Filename) }
func (f *File) Open(mode string) error {
	if f.Filename == "!STDIN!" {
		f.Reader = bufio.NewReader(os.Stdin)
		return nil
	}
	if f.Filename == "!STDOUT!" {
		f.Writer = bufio.NewWriter(os.Stdout)
		return nil
	}
	if f.Filename == "!STDERR!" {
		f.Writer = bufio.NewWriter(os.Stderr)
		return nil
	}

	md := os.O_RDONLY

	if mode == "w" {
		md = os.O_WRONLY

		os.Remove(f.Filename)
	} else {
		if strings.Contains(mode, "w") && strings.Contains(mode, "a") {
			md = os.O_WRONLY
			md |= os.O_APPEND
		}
	}

	file, err := os.OpenFile(f.Filename, os.O_CREATE|md, 0644)
	if err != nil {
		return err
	}

	f.Handle = file

	if md == os.O_RDONLY {
		f.Reader = bufio.NewReader(file)
	} else {
		f.Writer = bufio.NewWriter(file)
	}

	return nil
}

func (f *File) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "close" {
		f.Handle.Close()
		return &Boolean{Value: true}
	}
	if method == "lines" {

		// Do we not have a reader?
		if f.Reader == nil {
			return (&Null{})
		}

		// Result.
		var lines []string
		for {
			line, err := f.Reader.ReadString('\n')
			if err != nil {
				break
			}
			lines = append(lines, line)
		}

		// make results
		l := len(lines)
		result := make([]Object, l)
		for i, txt := range lines {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	if method == "methods" {
		static := []string{"methods"}
		dynamic := env.Names("file.")

		var names []string
		names = append(names, static...)
		for _, e := range dynamic {
			bits := strings.Split(e, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)

		result := make([]Object, len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	if method == "read" {
		// Check we have a reader.
		if f.Reader == nil {
			return (&String{Value: ""})
		}

		// Read and return a line.
		line, err := f.Reader.ReadString('\n')
		if err != nil {
			return (&String{Value: ""})
		}
		return &String{Value: line}
	}
	if method == "rewind" {
		// Rewind a handle by seeking to the beginning of the file.
		f.Handle.Seek(0, 0)
		return &Boolean{Value: true}
	}
	if method == "write" {

		// check we have an argument to write.
		if len(args) < 1 {
			return &Error{Message: "Missing argument to write()!"}
		}

		// Ensure we have a writer.
		if f.Writer == nil {
			return (&Null{})
		}

		// Write the text - coorcing to a string first.
		txt := args[0].Inspect()
		_, err := f.Writer.Write([]byte(txt))
		if err == nil {
			f.Writer.Flush()
			return &Boolean{Value: true}
		}

		return &Boolean{Value: false}
	}
	return nil
}
