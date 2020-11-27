package statistics

import (
	"testing"
)

func TestStandardDeviationShouldReturnExpectedResult(t *testing.T) {

	var samples = []float64{727.7, 1086.5, 1091.0, 1361.3, 1490.5, 1956.1}

	result, _ := StandardDeviation(samples)

	if !FloatEqualsWithEpsilon(result, 420.96, 0.003) {
		t.Errorf("Failed to caclulate standard deviation! Expected : 420.96 Result : %f ", result)
	}
}
