package object

var NIL = new(Nil)

type Nil struct{}

func (n *Nil) Type() ObjectType { return NIL_OBJ }
func (n *Nil) Inspect() string  { return "nil" }
func (n *Nil) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(n, method, env, args)
}

func (n *Nil) ToStringObj() *String {
	return NewString("")
}

func (n *Nil) ToIntegerObj() *Integer {
	return NewInteger(0.0)
}

func (n *Nil) ToFloatObj() *Float {
	return NewFloat(0.0)
}

func init() {
	objectMethods[NIL_OBJ] = map[string]ObjectMethod{}
}
