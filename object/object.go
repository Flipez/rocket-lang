package object

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/flipez/rocket-lang/ast"
)

type ObjectType string

var Evaluator func(node ast.Node, env *Environment) Object

type Object interface {
	Type() ObjectType
	Inspect() string
	InvokeMethod(method string, env Environment, args ...Object) Object
}

type Iterable interface {
	GetIterator(int, int, bool) Iterator
}

type Iterator interface {
	Next() (Object, Object, bool)
}

type Hashable interface {
	HashKey() HashKey
}

type Serializable interface {
	MarshalJSON() ([]byte, error)
}

type Stringable interface {
	ToStringObj() *String
}

// Integerable is implemented by objects that can attempt an integer
// conversion. The result is NIL when the value cannot be converted.
type Integerable interface {
	ToIntegerObj() Object
}

// Floatable is implemented by objects that can attempt a float
// conversion. The result is NIL when the value cannot be converted.
type Floatable interface {
	ToFloatObj() Object
}

// Type groups name a set of accepted types by what an object can do rather than
// by listing the types that can do it. They go wherever an ObjectType goes:
// Arg(HASHABLE), OptArg(ANY), OverloadArg(NUMERIC).
//
// Listing types by hand is what these replace. "Any value" used to be written
// four different ways, no two of them the same set: append accepted FUNCTION but
// not FLOAT, so [1].append(1.5) was an error, and include? accepted neither FLOAT
// nor MATRIX. A group is checked by asking the object, so a type added later
// joins without anyone remembering to update a list.
const (
	// ANY accepts every object.
	ANY = "ANY"
	// HASHABLE accepts what can be a hash key, which is what a Hash's key
	// arguments need. Declaring these ANY let a NIL or a MATRIX through to an
	// unchecked type assertion, and {"a": 1}.get(nil, 0) panicked.
	HASHABLE = "HASHABLE"
	// STRINGABLE accepts what has a string form, which is what join needs of
	// every element it is given.
	STRINGABLE = "STRINGABLE"
	// INTEGERABLE accepts what can be read as an integer. Wider than NUMERIC: a
	// string that parses and a boolean qualify, which is why ["12"].sum() is 12.
	INTEGERABLE = "INTEGERABLE"
	// COMPARABLE accepts what can be ordered against its own kind. Ordering
	// also requires the values to be of one type, which is a property of the
	// collection rather than of any single value, so sort checks that on top.
	COMPARABLE = "COMPARABLE"
	// NUMERIC accepts INTEGER and FLOAT.
	NUMERIC = "NUMERIC"
	// CALLABLE accepts what can be called: a function written in RocketLang, or
	// a builtin such as puts, which is a value too.
	CALLABLE = "CALLABLE"
)

// typeGroups maps each group to the question it asks of an object. Where a Go
// interface already expresses the requirement, the group asserts against it, so
// the group and the method body cannot disagree about what qualifies.
var typeGroups = map[string]func(Object) bool{
	ANY: func(Object) bool { return true },
	HASHABLE: func(o Object) bool {
		_, ok := o.(Hashable)
		return ok
	},
	STRINGABLE: func(o Object) bool {
		_, ok := o.(Stringable)
		return ok
	},
	INTEGERABLE: func(o Object) bool {
		_, ok := o.(Integerable)
		return ok
	},
	COMPARABLE: func(o Object) bool {
		switch o.Type() {
		case INTEGER_OBJ, FLOAT_OBJ, STRING_OBJ:
			return true
		default:
			return false
		}
	},
	NUMERIC: func(o Object) bool {
		return o.Type() == INTEGER_OBJ || o.Type() == FLOAT_OBJ
	},
	CALLABLE: func(o Object) bool {
		return o.Type() == FUNCTION_OBJ || o.Type() == BUILTIN_FUNCTION_OBJ
	},
}

const (
	INTEGER_OBJ          = "INTEGER"
	FLOAT_OBJ            = "FLOAT"
	BOOLEAN_OBJ          = "BOOLEAN"
	NIL_OBJ              = "NIL"
	RETURN_VALUE_OBJ     = "RETURN_VALUE"
	BREAK_VALUE_OBJ      = "BREAK_VALUE"
	NEXT_VALUE_OBJ       = "NEXT_VALUE"
	ERROR_OBJ            = "ERROR"
	FUNCTION_OBJ         = "FUNCTION"
	STRING_OBJ           = "STRING"
	ARRAY_OBJ            = "ARRAY"
	HASH_OBJ             = "HASH"
	MATRIX_OBJ           = "MATRIX"
	FILE_OBJ             = "FILE"
	MODULE_OBJ           = "MODULE"
	HTTP_OBJ             = "HTTP"
	BUILTIN_MODULE_OBJ   = "BUILTIN_MODULE"
	BUILTIN_FUNCTION_OBJ = "BUILTIN_FUNCTION"
	BUILTIN_PROPERTY_OBJ = "BUILTIN_PROPERTY"
)

// knownObjectTypes is every name a type can go by, which is what lets is_a?
// tell an unknown name from one that simply does not match. A new type belongs
// here; TestKnownObjectTypesAreComplete checks the ones that register methods.
var knownObjectTypes = map[string]bool{
	INTEGER_OBJ:          true,
	FLOAT_OBJ:            true,
	BOOLEAN_OBJ:          true,
	NIL_OBJ:              true,
	RETURN_VALUE_OBJ:     true,
	BREAK_VALUE_OBJ:      true,
	NEXT_VALUE_OBJ:       true,
	ERROR_OBJ:            true,
	FUNCTION_OBJ:         true,
	STRING_OBJ:           true,
	ARRAY_OBJ:            true,
	HASH_OBJ:             true,
	MATRIX_OBJ:           true,
	FILE_OBJ:             true,
	MODULE_OBJ:           true,
	HTTP_OBJ:             true,
	BUILTIN_MODULE_OBJ:   true,
	BUILTIN_FUNCTION_OBJ: true,
	BUILTIN_PROPERTY_OBJ: true,
}

// KnownObjectTypes returns the type names, sorted. Exported for the tests that
// check the type groups and is_a? against every type there is.
func KnownObjectTypes() []string {
	names := make([]string, 0, len(knownObjectTypes))
	for name := range knownObjectTypes {
		names = append(names, name)
	}
	sort.Strings(names)

	return names
}

// TypeGroupNames returns the group names, sorted.
func TypeGroupNames() []string {
	names := make([]string, 0, len(typeGroups))
	for name := range typeGroups {
		names = append(names, name)
	}
	sort.Strings(names)

	return names
}

type Argument struct {
	Types    []string
	Optional bool
	Overload bool
}

func (a Argument) String() string {
	return strings.Join(a.Types, "|")
}

// usage renders the argument as it appears in a signature. Brackets mean it may
// be left out, a trailing ellipsis means more of the same may follow. Without
// them count(STRING), split(STRING) and starts_with?(STRING) all printed
// identically while meaning "exactly one", "zero or one" and "one or more".
//
// This is deliberately not String(), which is used in the "wrong argument type"
// error, where the reader is being told what a single position accepts and the
// brackets would only be noise.
func (a Argument) usage() string {
	rendered := strings.Join(a.Types, "|")

	if a.Overload {
		// Parenthesise a union first, so the ellipsis reads as applying to the
		// whole argument rather than only to the last type of it.
		if len(a.Types) > 1 {
			rendered = "(" + rendered + ")"
		}
		rendered += "..."
	}

	if a.Optional {
		rendered = "[" + rendered + "]"
	}

	return rendered
}

// InGroup reports whether o belongs to the named group. Argument patterns get
// this through Check; method bodies need it directly, because a requirement on
// the elements of a collection is not something an argument pattern can state:
// join needs every element STRINGABLE, not its separator.
func InGroup(group string, o Object) bool {
	check, ok := typeGroups[group]
	if !ok {
		return false
	}

	return check(o)
}

func (a Argument) Check(o Object) bool {
	for _, t := range a.Types {
		if group, isGroup := typeGroups[t]; isGroup {
			if group(o) {
				return true
			}

			continue
		}

		if ObjectType(t) == o.Type() {
			return true
		}
	}

	return false
}

func Arg(types ...string) Argument {
	return Argument{Types: types}
}

func OptArg(types ...string) Argument {
	return Argument{Types: types, Optional: true}
}

func OverloadArg(types ...string) Argument {
	return Argument{Types: types, Overload: true}
}

func Args(args ...Argument) []Argument {
	return args
}

type MethodLayout struct {
	ArgPattern    []Argument
	ReturnPattern []Argument
	Description   string
	Input         string
	Output        string
}

func (ml MethodLayout) requiredArgs() []Argument {
	args := make([]Argument, 0)
	for _, arg := range ml.ArgPattern {
		if arg.Optional {
			continue
		}
		args = append(args, arg)
	}
	return args
}

type ObjectMethod struct {
	Layout MethodLayout
	method func(Object, []Object, Environment) Object
}

func (ml MethodLayout) validateArgs(args []Object) error {
	requiredArgs := ml.requiredArgs()
	requiredArgsLen := len(requiredArgs)
	possibleArgsLen := len(ml.ArgPattern)
	givenArgsLen := len(args)
	overloadArgs := possibleArgsLen > 0 && ml.ArgPattern[possibleArgsLen-1].Overload

	if givenArgsLen < requiredArgsLen {
		return fmt.Errorf("too few arguments: got=%d, want=%d", givenArgsLen, requiredArgsLen)
	}

	if (givenArgsLen > possibleArgsLen) && !overloadArgs {
		return fmt.Errorf("too many arguments: got=%d, want=%d", givenArgsLen, possibleArgsLen)
	}

	if givenArgsLen > 0 {
		for idx, arg := range args {
			var argToCheck Argument
			if (idx >= possibleArgsLen) && overloadArgs {
				argToCheck = ml.ArgPattern[possibleArgsLen-1]
			} else {
				argToCheck = ml.ArgPattern[idx]
			}

			if !argToCheck.Check(arg) {
				return fmt.Errorf("wrong argument type on position %d: got=%s, want=%s", idx+1, arg.Type(), argToCheck)
			}
		}
	}

	return nil
}

func (ml MethodLayout) DocsReturnPattern() string {
	types := make([]string, len(ml.ReturnPattern))
	for idx, pattern := range ml.ReturnPattern {
		types[idx] = strings.Join(pattern.Types, "|")
	}
	return strings.Join(types, ", ")
}

func (ml MethodLayout) Usage(name string) string {
	var args string

	if len(ml.ArgPattern) > 0 {
		types := make([]string, len(ml.ArgPattern))
		for idx, pattern := range ml.ArgPattern {
			types[idx] = pattern.usage()
		}
		args = strings.Join(types, ", ")
	}

	return fmt.Sprintf("%s(%s)", name, args)
}

func (om ObjectMethod) Call(o Object, args []Object, env Environment) Object {
	if err := om.Layout.validateArgs(args); err != nil {
		return NewError(err)
	}
	return om.method(o, args, env)
}

var objectMethods = make(map[ObjectType]map[string]ObjectMethod)

func ListObjectMethods() map[ObjectType]map[string]ObjectMethod {
	return objectMethods
}

func init() {
	objectMethods["*"] = map[string]ObjectMethod{
		"to_string": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				if stringable, ok := o.(Stringable); ok {
					return stringable.ToStringObj()
				}

				return NewString("")
			},
		},
		"to_integer": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ, NIL_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				if integerable, ok := o.(Integerable); ok {
					return integerable.ToIntegerObj()
				}

				return NIL
			},
		},
		"to_float": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(FLOAT_OBJ, NIL_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				if floatable, ok := o.(Floatable); ok {
					return floatable.ToFloatObj()
				}

				return NIL
			},
		},
		"to_json": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(STRING_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				if serializeableObject, ok := o.(Serializable); ok {
					j, err := json.Marshal(serializeableObject)
					if err != nil {
						return NewErrorFormat("Error while marshal value: %s", err.Error())
					}
					return NewString(string(j))
				}

				return NewErrorFormat("%s is not serializable", o.Type())
			},
		},
		"is_a?": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(STRING_OBJ),
				),
				ReturnPattern: Args(
					Arg(BOOLEAN_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				name := args[0].(*String).Value

				// A name that is neither is a mistake rather than a "no": a
				// typo would otherwise answer false and read as a real result.
				if !knownObjectTypes[name] {
					if _, isGroup := typeGroups[name]; !isGroup {
						return NewErrorFormat("unknown type or type group: %s", name)
					}
				}

				// The same resolution the argument patterns use, so what a
				// value says about itself and what a method accepts can never
				// drift apart.
				if Arg(name).Check(o) {
					return TRUE
				}

				return FALSE
			},
		},
		"type_groups": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				// Sorted, for the same reason methods() is: read out of the map
				// the order differed between runs.
				groups := make([]Object, 0, len(typeGroups))
				for _, name := range TypeGroupNames() {
					// ANY is left out. Every other group says something the
					// value can do; ANY says an argument accepts anything,
					// which is a statement about a parameter rather than about
					// this value. As a property it is true of everything, so it
					// distinguishes nothing and only prefixes every list with a
					// constant. is_a?("ANY") still answers true, because ANY is
					// a real group wherever a signature uses it.
					if name == ANY {
						continue
					}

					if typeGroups[name](o) {
						groups = append(groups, NewString(name))
					}
				}

				return NewArray(groups)
			},
		},
		"nil?": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(BOOLEAN_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				if o.Type() == NIL_OBJ {
					return TRUE
				}

				return FALSE
			},
		},
		"methods": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				oms := objectMethods[o.Type()]
				names := make([]string, 0, len(oms))
				for name := range oms {
					names = append(names, name)
				}
				// Sorted, because reading it straight out of the map made the
				// order differ between runs of the same program.
				sort.Strings(names)

				result := make([]Object, len(names))
				for i, name := range names {
					result[i] = NewString(name)
				}

				return NewArray(result)
			},
		},
		"type": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return NewString(string(o.Type()))
			},
		},
	}

	// help prints the receiver's method list. It is printed rather than
	// returned because this listing exists to be read, and Inspect() would
	// escape the newlines onto one line.
	helpMethod := ObjectMethod{
		Layout: MethodLayout{
			ReturnPattern: Args(
				Arg(NIL_OBJ),
			),
		},
		method: func(o Object, _ []Object, _ Environment) Object {
			oms := objectMethods[o.Type()]
			names := make([]string, 0, len(oms))
			for name := range oms {
				names = append(names, name)
			}
			sort.Strings(names)

			fmt.Printf("%s supports the following methods:\n", o.Type())
			for _, name := range names {
				fmt.Printf("\t%s\n", oms[name].Layout.Usage(name))
			}

			return NIL
		},
	}

	objectMethods["*"]["help"] = helpMethod
	// wat is the only alias in the language: a deliberate easter egg, kept
	// because it predates help. It is not a precedent -- every other method
	// has exactly one name.
	objectMethods["*"]["wat"] = helpMethod
}

func objectMethodLookup(o Object, method string, env Environment, args []Object) Object {
	if oms, ok := objectMethods[o.Type()]; ok {
		if objMethod, ok := oms[method]; ok {
			return objMethod.Call(o, args, env)
		}
	}

	if oms, ok := objectMethods["*"]; ok {
		if objMethod, ok := oms[method]; ok {
			return objMethod.Call(o, args, env)
		}
	}

	return nil
}

func CompareObjects(ao, bo Object) bool {
	switch ao.Type() {
	case NIL_OBJ:
		return bo.Type() == NIL_OBJ
	case INTEGER_OBJ:
		if b, ok := bo.(*Integer); ok {
			return ao.(*Integer).Value == b.Value
		}
		return false
	case FLOAT_OBJ:
		if b, ok := bo.(*Float); ok {
			return ao.(*Float).Value == b.Value
		}
		return false
	case BOOLEAN_OBJ:
		if b, ok := bo.(*Boolean); ok {
			return ao.(*Boolean).Value == b.Value
		}
		return false
	case ERROR_OBJ:
		if b, ok := bo.(*Error); ok {
			return ao.(*Error).Message == b.Message
		}
		return false
	case STRING_OBJ:
		if b, ok := bo.(*String); ok {
			return ao.(*String).Value == b.Value
		}
		return false
	case ARRAY_OBJ:
		if b, ok := bo.(*Array); ok {
			a, _ := ao.(*Array)

			if len(a.Elements) != len(b.Elements) {
				return false
			}

			for idx, element := range a.Elements {
				if !CompareObjects(element, b.Elements[idx]) {
					return false
				}
			}

			return true
		}
		return false
	case HASH_OBJ:
		if b, ok := bo.(*Hash); ok {
			a, _ := ao.(*Hash)

			if len(a.Pairs) != len(b.Pairs) {
				return false
			}

			for aKey, aPair := range a.Pairs {
				bPair, ok := b.Pairs[aKey]
				if !ok {
					return false
				}
				if !CompareObjects(aPair.Key, bPair.Key) {
					return false
				}
				if !CompareObjects(aPair.Value, bPair.Value) {
					return false
				}
			}

			return true
		}
		return false
	}

	return false
}

func IsError(o Object) bool {
	return o != nil && o.Type() == ERROR_OBJ
}

func IsNumber(o Object) bool {
	return o != nil && (o.Type() == INTEGER_OBJ || o.Type() == FLOAT_OBJ)
}

func IsTruthy(o Object) bool {
	switch o {
	case NIL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func IsFalsy(o Object) bool {
	return !IsTruthy(o)
}

func AddEvaluator(e func(node ast.Node, env *Environment) Object) {
	Evaluator = e
}

func AnyToObject(a any) Object {
	switch v := a.(type) {
	case bool:
		if v {
			return TRUE
		}
		return FALSE
	case string:
		return NewString(v)
	case int:
		return NewInteger(v)
	case int64:
		return NewInteger(int(v))
	case float64:
		return NewFloat(v)
	case []any:
		arr := make([]Object, len(v))
		for idx, item := range v {
			arr[idx] = AnyToObject(item)
		}
		return NewArray(arr)
	case map[string]any:
		hash := NewHash(nil)
		for key, val := range v {
			hash.Set(AnyToObject(key), AnyToObject(val))
		}
		return hash
	}
	return NIL
}

func ObjectToAny(o Object) any {
	switch v := o.(type) {
	case *Boolean:
		return v.Value
	case *String:
		return v.Value
	case *Integer:
		return v.Value
	case *Float:
		return v.Value
	case *Array:
		array := make([]any, len(v.Elements))
		for idx, element := range v.Elements {
			array[idx] = ObjectToAny(element)
		}
		return array
	case *Hash:
		hash := make(map[any]any)
		for _, pair := range v.Pairs {
			key := ObjectToAny(pair.Key)
			if key != nil {
				hash[key] = ObjectToAny(pair.Value)
			}
		}
		return hash
	}
	return nil
}
