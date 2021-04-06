package algebra

import "fmt"

func MultiplyMatrices(matrixOne [][]float64, matrixTwo [][]float64) (resultMatrix [][]float64, err error) {

	if matrixOne == nil {
		return nil, fmt.Errorf("MultiplyMatrices : invalid matrixOne")
	}

	if len(matrixOne) == 0 || matrixOne[0] == nil {
		return nil, fmt.Errorf("MultiplyMatrices : invalid rows of matrixOne")
	}

	if matrixTwo == nil {
		return nil, fmt.Errorf("MultiplyMatrices : invalid matrixTwo")
	}

	if len(matrixTwo) == 0 || matrixTwo[0] == nil {
		return nil, fmt.Errorf("MultiplyMatrices : invalid rows of matrixTwo")
	}

	rows := len(matrixOne[0])
	columns := len(matrixTwo)

	hasNoValidMatricesSizeForMultiplication := rows != columns
	if hasNoValidMatricesSizeForMultiplication {
		return nil, fmt.Errorf(" MultiplyMatrices : number of rows for the first matrix is different from number of columns of the second matrix ")
	}

	rowsOfMatrixOne := len(matrixOne)
	columnsOfMatrixOne := len(matrixOne[0])

	rowsOfMatrixTwo := len(matrixTwo)
	columnsOfMatrixTwo := len(matrixTwo[0])

	resultMatrix = make([][]float64, rowsOfMatrixOne)

	for rOne := 0; rOne < rowsOfMatrixOne; rOne++ {
		for cOne := 0; cOne < columnsOfMatrixOne; cOne++ {
			for rTwo := 0; rTwo < rowsOfMatrixTwo; rTwo++ {
				resultMatrix[rOne] = make([]float64, columnsOfMatrixTwo)
				for cTwo := 0; cTwo < columnsOfMatrixTwo; cTwo++ {
					for multiplicationIndex := 0; multiplicationIndex < columnsOfMatrixOne; multiplicationIndex++ {
						resultMatrix[rOne][cTwo] += matrixOne[rOne][multiplicationIndex] * matrixTwo[multiplicationIndex][cTwo]
					}
				}
			}
		}
	}

	return resultMatrix, nil
}
