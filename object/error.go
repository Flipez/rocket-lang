package object

import "fmt"

type Error struct {
	Message string
}

func NewError(e interface{}) *Error {
	return &Error{Message: fmt.Sprintf("%s", e)}
}

func NewErrorFormat(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
func (e *Error) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(e, method, env, args)
}

func init() {
	objectMethods[ERROR_OBJ] = map[string]ObjectMethod{
		"msg": ObjectMethod{
			Layout: MethodLayout{
				Description: "Returns the error message",
				Example: `» def ()
puts(nope)
rescue
puts((rescued error: + error.msg()))
end
🚀 » test()
"rescued error:identifier not found: nope"`,
				ReturnPattern: Args(
					Arg(STRING_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				e := o.(*Error)
				return NewString(e.Message)
			},
		},
	}
}
