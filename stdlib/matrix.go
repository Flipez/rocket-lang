package stdlib

import (
	"github.com/flipez/rocket-lang/object"
)

func init() {
	RegisterFunction("Matrix",
		object.MethodLayout{
			ArgPattern:    object.Args(object.Arg(object.ARRAY_OBJ)),
			ReturnPattern: object.Args(object.Arg(object.MATRIX_OBJ, object.ERROR_OBJ)),
		},
		matrixConstructor)
}

func matrixConstructor(_ object.Environment, args ...object.Object) object.Object {
	arr := args[0].(*object.Array)
	matrix, err := object.NewMatrixFromObjects(arr.Elements)
	if err != nil {
		return object.NewErrorFormat("failed to create matrix: %s", err.Error())
	}
	return matrix
}
