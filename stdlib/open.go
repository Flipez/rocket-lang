package stdlib

import (
	"github.com/flipez/rocket-lang/object"
)

func openFunction(env object.Environment, args ...object.Object) object.Object {
  mode := "r"
  perm := "0644"

	path := args[0].(*object.String).Value
  if len(args) > 1 {
	  mode = args[1].(*object.String).Value
  }
  if len(args) > 2 {
	  perm = args[2].(*object.String).Value
  }

	file := object.NewFile(path)
	err := file.Open(mode, perm)
	if err != nil {
		return object.NewErrorFormat(err.Error())
	}
	return (file)
}
