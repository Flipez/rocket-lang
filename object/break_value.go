package object

// BreakValue signals that a `break` was reached. It carries no payload:
// `break` took a value until 0.18, when that form was removed, so the loops
// have nothing to hand back.
type BreakValue struct{}

func NewBreakValue() *BreakValue {
	return &BreakValue{}
}

func (bv *BreakValue) Type() ObjectType { return BREAK_VALUE_OBJ }
func (bv *BreakValue) Inspect() string  { return "break" }
func (bv *BreakValue) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(bv, method, env, args)
}
