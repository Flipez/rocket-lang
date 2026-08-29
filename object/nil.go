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

// MarshalJSON renders nil as JSON null. Without it Nil was not Serializable at
// all, so nil.to_json() was an error, and a nil inside an array came out as {}
// because json.Marshal fell back to encoding the empty struct.
func (n *Nil) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

func init() {
	objectMethods[NIL_OBJ] = map[string]ObjectMethod{}
}
