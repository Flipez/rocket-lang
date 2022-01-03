package object

import "fmt"

type Module struct {
	Name       string
	Attributes Object
}

func (m *Module) Type() ObjectType { return MODULE_OBJ }
func (m *Module) Inspect() string  { return fmt.Sprintf("module(%s)", m.Name) }
func (m *Module) InvokeMethod(method string, env Environment, args ...Object) Object {
	return nil
}
