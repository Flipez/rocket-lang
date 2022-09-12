package object

var NIL = new(Nil)

type Nil struct{}

func (n *Nil) Type() ObjectType { return NIL_OBJ }
func (n *Nil) Inspect() string  { return "nil" }
func (n *Nil) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(n, method, env, args)
}
func init() {
	objectMethods[NIL_OBJ] = map[string]ObjectMethod{
		"plz_s": ObjectMethod{
			Layout: MethodLayout{
				Description: "Returns empty string.",
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(_ Object, _ []Object, _ Environment) Object {
				return NewString("")
			},
		},
		"plz_i": ObjectMethod{
			Layout: MethodLayout{
				Description: "Returns zero integer.",
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(_ Object, _ []Object, _ Environment) Object {
				return NewInteger(0)
			},
		},
		"plz_f": ObjectMethod{
			Layout: MethodLayout{
				Description: "Returns zero float.",
				ReturnPattern: Args(
					Arg(FLOAT_OBJ),
				),
			},
			method: func(_ Object, _ []Object, _ Environment) Object {
				return NewFloat(0)
			},
		},
	}
}
