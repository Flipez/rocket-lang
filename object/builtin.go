package object

type BuiltinModule struct {
	Name       string
	Functions  map[string]*BuiltinFunction
	Properties map[string]*BuiltinProperty
}

func NewBuiltinModule(name string, funcs map[string]*BuiltinFunction, props map[string]*BuiltinProperty) *BuiltinModule {
	return &BuiltinModule{
		Name:       name,
		Functions:  funcs,
		Properties: props,
	}
}

func (b *BuiltinModule) Type() ObjectType { return BUILTIN_MODULE_OBJ }
func (b *BuiltinModule) Inspect() string  { return b.Name }
func (b *BuiltinModule) InvokeMethod(method string, env Environment, args ...Object) Object {
	if fn, ok := b.Functions[method]; ok {
		return fn.Fn(env, args...)
	}

	return objectMethodLookup(b, method, env, args)
}

type BuiltinFunction struct {
	Name string
	Fn   func(Environment, ...Object) Object
}

func NewBuiltinFunction(name string, f func(Environment, ...Object) Object) *BuiltinFunction {
	return &BuiltinFunction{
		Name: name,
		Fn:   f,
	}
}

func (b *BuiltinFunction) Type() ObjectType { return BUILTIN_FUNCTION_OBJ }
func (b *BuiltinFunction) Inspect() string  { return b.Name }
func (b *BuiltinFunction) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(b, method, env, args)
}

type BuiltinProperty struct {
	Name  string
	Value Object
}

func NewBuiltinProperty(name string, val Object) *BuiltinProperty {
	return &BuiltinProperty{
		Name:  name,
		Value: val,
	}
}

func (b *BuiltinProperty) Type() ObjectType { return BUILTIN_PROPERTY_OBJ }
func (b *BuiltinProperty) Inspect() string  { return b.Name }
func (b *BuiltinProperty) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(b, method, env, args)
}
