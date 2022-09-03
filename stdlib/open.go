package stdlib

import (
	"github.com/flipez/rocket-lang/object"
)

func openFunction(env object.Environment, args ...object.Object) object.Object {
	path := args[0].(*object.String).Value
	mode := args[1].(*object.String).Value
	perm := args[2].(*object.String).Value

	file := object.NewFile(path)
	err := file.Open(mode, perm)
	if err != nil {
		return object.NewErrorFormat(err.Error())
	}
	return (file)
}
