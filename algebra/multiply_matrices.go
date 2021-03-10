package algebra

import "errors"

func MultiplyMatrices(matrixOne [][]float64, matrixTwo [][]float64) (resultMatrix [][]float64, err error) {

	if len(matrixOne) != len(matrixTwo) {
		return nil, errors.New("cannot multiply the given matrices!")
	}

	resultMatrix = make([][]float64, len(matrixOne))

	for i := 0; i < len(matrixOne); i++ {
		resultMatrix[i] = make([]float64, len(matrixTwo))
		for j := 0; j < len(matrixTwo); j++ {
			for k := 0; k < len(matrixTwo); k++ {
				resultMatrix[i][j] += matrixOne[i][k] * matrixTwo[k][j]
			}
		}
	}

	return resultMatrix, nil
}
