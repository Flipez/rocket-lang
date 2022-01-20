package object

var NULL = new(Null)

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }
func (n *Null) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(n, method, args)
}
func init() {
	objectMethods[NULL_OBJ] = map[string]ObjectMethod{
		"plz_s": ObjectMethod{
			description: "Returns empty string.",
			returnPattern: [][]string{
				[]string{STRING_OBJ},
			},
			method: func(_ Object, _ []Object) Object {
				return NewString("")
			},
		},
		"plz_i": ObjectMethod{
			description: "Returns zero integer.",
			returnPattern: [][]string{
				[]string{INTEGER_OBJ},
			},
			method: func(_ Object, _ []Object) Object {
				return NewInteger(0)
			},
		},
		"plz_f": ObjectMethod{
			description: "Returns zero float.",
			returnPattern: [][]string{
				[]string{FLOAT_OBJ},
			},
			method: func(_ Object, _ []Object) Object {
				return NewFloat(0)
			},
		},
	}
}
