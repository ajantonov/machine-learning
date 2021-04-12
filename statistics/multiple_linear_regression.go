package statistics

import "github.com/ajantonov/machine-learning/algebra"

/* MultipleRegression - calculates beta coefficients for multiple predictors and response y
 */
func MultipleRegression(matrixX [][]float64, matrixY [][]float64) ([][]float64, error) {
	matrixTransposedOfX, err := algebra.TransposeMatrix(matrixX)
	if err != nil {
		return nil, err
	}

	matrixTransposedOfXByX, err := algebra.MultiplyMatrices(matrixTransposedOfX, matrixX)
	if err != nil {
		return nil, err
	}

	matrixInversedOfTransposedOfXByX, err := algebra.InverseMatrix(matrixTransposedOfXByX)
	if err != nil {
		return nil, err
	}

	matrixHat, err := algebra.MultiplyMatrices(matrixInversedOfTransposedOfXByX, matrixTransposedOfX)
	if err != nil {
		return nil, err
	}

	matrixResult, err := algebra.MultiplyMatrices(matrixHat, matrixY)
	if err != nil {
		return nil, err
	}

	return matrixResult, nil
}
