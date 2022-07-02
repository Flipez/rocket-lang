package object

type BreakValue struct {
	Value Object
}

func NewBreakValue(o Object) *BreakValue {
	return &BreakValue{Value: o}
}

func (bv *BreakValue) Type() ObjectType { return BREAK_VALUE_OBJ }
func (bv *BreakValue) Inspect() string  { return bv.Value.Inspect() }
func (bv *BreakValue) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(bv, method, env, args)
}
