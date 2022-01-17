package object

type ReturnValue struct {
	Value Object
}

func NewReturnValue(o Object) *ReturnValue {
	return &ReturnValue{Value: o}
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
func (rv *ReturnValue) InvokeMethod(method string, env Environment, args ...Object) Object {
	return nil
}
