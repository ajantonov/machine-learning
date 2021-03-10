package algebra

import (
	"math"
)

const EPSILON float64 = 0.0001

/*
	FloatEquals - unit test helper function
*/
func FloatEquals(leftVariable, rightVariable float64) bool {
	if math.Abs(leftVariable-rightVariable) < EPSILON {
		println(math.Abs(leftVariable - rightVariable))
		return true
	}
	return false
}

func FloatEqualsWithEpsilon(leftVariable, rightVariable float64, epsilon float64) bool {
	if math.Abs(leftVariable-rightVariable) < epsilon {
		println(math.Abs(leftVariable - rightVariable))
		return true
	}
	return false
}
