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
			description: "Returns empty string.",
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(_ Object, _ []Object, _ Environment) Object {
				return NewString("")
			},
		},
		"plz_i": ObjectMethod{
			description: "Returns zero integer.",
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			method: func(_ Object, _ []Object, _ Environment) Object {
				return NewInteger(0)
			},
		},
		"plz_f": ObjectMethod{
			description: "Returns zero float.",
			returnPattern: [][]string{
				[]string{FLOAT_OBJ},
			},
			method: func(_ Object, _ []Object, _ Environment) Object {
				return NewFloat(0)
			},
		},
	}
}
