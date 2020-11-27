package statistics

import (
	"math"
)

const EPSILON float64 = 0.0001

/*
	FloatEquals - unit test helper function
*/
func FloatEquals(a, b float64) bool {
	if math.Abs(a-b) < EPSILON {
		println(math.Abs(a - b))
		return true
	}
	return false
}

func FloatEqualsWithEpsilon(a, b float64, epsilon float64) bool {
	if math.Abs(a-b) < epsilon {
		println(math.Abs(a - b))
		return true
	}
	return false
}
