package statistics

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

const EPSILON float64 = 0.0001

/*
	FloatEquals - unit test helper function
*/
func FloatEquals(leftVariable, rightVariable float64) bool {
	if math.Abs(leftVariable-rightVariable) < EPSILON {
		math.Abs(leftVariable - rightVariable)
		return true
	}
	return false
}

func FloatEqualsWithEpsilon(leftVariable, rightVariable float64, epsilon float64) bool {
	if math.Abs(leftVariable-rightVariable) < epsilon {
		math.Abs(leftVariable - rightVariable)
		return true
	}
	return false
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
