package algebra

import "fmt"

func MultiplyMatrices(matrixOne [][]float64, matrixTwo [][]float64) (resultMatrix [][]float64, err error) {

	if matrixOne == nil {
		return nil, fmt.Errorf("Invalid matrixOne")
	}

	if len(matrixOne) == 0 || matrixOne[0] == nil {
		return nil, fmt.Errorf("Invalid rows of matrixOne")
	}

	if len(matrixTwo) == 0 || matrixTwo[0] == nil {
		return nil, fmt.Errorf("Invalid matrixTwo")
	}

	rows := len(matrixOne)
	columns := len(matrixTwo[0])

	hasNoValidMatricesSizeForMultiplication := rows != columns
	if hasNoValidMatricesSizeForMultiplication {
		return nil, fmt.Errorf(" MultiplyMatrices : number of rows for the first matrix is different from number of columns of the second matrix ")
	}

	//resultMatrix = make([][]float64, len(matrixOne))
	//for i, x := range
	//
	return resultMatrix, nil
}
