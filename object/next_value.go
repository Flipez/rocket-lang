package object

type NextValue struct {
	Value Object
}

func NewNextValue(o Object) *NextValue {
	return &NextValue{Value: o}
}

func (nv *NextValue) Type() ObjectType { return NEXT_VALUE_OBJ }
func (nv *NextValue) Inspect() string  { return nv.Value.Inspect() }
func (nv *NextValue) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(nv, method, env, args)
}
