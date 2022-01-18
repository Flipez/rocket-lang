package object

type Builtin struct {
	Name string
	Fn   BuiltinFunction
}

func NewBuiltin(name string, f BuiltinFunction) *Builtin {
	return &Builtin{
		Name: name,
		Fn:   f,
	}
}

type BuiltinFunction func(args ...Object) Object

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }
func (b *Builtin) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(b, method, args)
}
