package statistics

import "errors"

func LinearRegression(inputs []float64, outputs []float64) (alpha float64, beta float64, err error) {

	if 0 == len(inputs) {
		return float64(0), float64(0), errors.New("invalid parameter inputs")
	}

	if 0 == len(outputs) {
		return float64(0), float64(0), errors.New("invalid parameter outputs")
	}

	// Calculate mean of inputs
	meanOfInputs, err := Mean(inputs)
	if nil != err {
		return float64(0), float64(0), errors.New("failed to calculate mean for inputs")
	}

	// Calculate mean of outputs
	meanOfOutputs, err := Mean(inputs)
	if nil != err {
		return float64(0), float64(0), errors.New("failed to calculate mean for outputs")
	}

	var numerator = float64(0)
	var denominator = float64(0)

	// Calculate alpha
	for index := 0; index < len(inputs); index++ {
		numerator += (inputs[index] - meanOfInputs) * (outputs[index] - meanOfOutputs)
		denominator += (inputs[index] - meanOfInputs) * (inputs[index] - meanOfInputs)
	}

	alpha = numerator / denominator

	// Calculate beta

	return alpha, float64(0), nil
}
