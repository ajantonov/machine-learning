package statistics

import "errors"

func LinearRegression(inputs []float64, outputs []float64) (alpha float64, beta float64, err error) {

	if 0 == len(inputs) {
		return float64(0), float64(0), errors.New("invalid parameter inputs")
	}

	if 0 == len(outputs) {
		return float64(0), float64(0), errors.New("invalid parameter outputs")
	}

	return float64(0), float64(0), nil
}
