package object

type Builtin struct {
	Name string
	Fn   BuiltinFunction
}

type BuiltinFunction func(args ...Object) Object

func (b *Builtin) Type() ObjectType                                                   { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string                                                    { return "builtin function" }
func (b *Builtin) InvokeMethod(method string, env Environment, args ...Object) Object { return nil }
