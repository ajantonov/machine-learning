package algebra

import (
	"fmt"
)

func MultiplyMatrices(matrixOne [][]float64, matrixTwo [][]float64) (resultMatrix [][]float64, err error) {

	if len(matrixOne) != len(matrixTwo[0]) {
		return nil,
			fmt.Errorf(" MultiplyMatrices : cannot multiply the given matrices rows %d != columns %d ",
				len(matrixOne), len(matrixTwo))
	}

	resultMatrix = make([][]float64, len(matrixOne))

	for i := 0; i < len(matrixOne); i++ {
		resultMatrix[i] = make([]float64, len(matrixTwo[i]))
		for j := 0; j < len(matrixTwo[i]); j++ {
			for k := 0; k < len(matrixTwo); k++ {
				resultMatrix[i][j] += matrixOne[i][k] * matrixTwo[k][j]
			}
		}
	}

	return resultMatrix, nil
}
