package statistics

import (
	"math"
)

const EPSILON float64 = 0.000001

func FloatEquals(a, b float64) bool {
	if math.Abs(a-b) < EPSILON {
		println(math.Abs(a - b))
		return true
	}
	return false
}