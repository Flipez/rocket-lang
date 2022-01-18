package object

var NULL = new(Null)

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }
func (n *Null) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(n, method, args)
}
