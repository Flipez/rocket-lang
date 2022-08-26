package stdlib

import (
	"math"

	"github.com/flipez/rocket-lang/object"
)

func registerConstants() {
	Constants["E"] = object.NewFloat(math.E)
	Constants["Pi"] = object.NewFloat(math.Pi)
	Constants["Phi"] = object.NewFloat(math.Phi)

	Constants["Sqrt2"] = object.NewFloat(math.Sqrt2)
	Constants["SqrtE"] = object.NewFloat(math.SqrtE)
	Constants["SqrtPi"] = object.NewFloat(math.SqrtPi)
	Constants["SqrtPhi"] = object.NewFloat(math.SqrtPhi)

	Constants["Ln2"] = object.NewFloat(math.Ln2)
	Constants["Log2E"] = object.NewFloat(math.Log2E)
	Constants["Ln10"] = object.NewFloat(math.Ln10)
	Constants["Log10E"] = object.NewFloat(math.Log10E)
}
