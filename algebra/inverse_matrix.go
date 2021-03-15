package algebra

import "errors"

func InverseMatrix(matrix [][]float64) ([][]float64, error) {
	size := len(matrix)

	var inverseMatrix = CreateMatrix(size, size)
	determinant := Determinant(matrix, size)
	if determinant == 0 {
		return nil, errors.New("Singular matrix, cannot find its inverse!")
	}

	adjoinMatrix, err := AdjointMatrix(matrix)
	if err != nil {
		return nil, err
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			inverseMatrix[i][j] = adjoinMatrix[i][j] / determinant
		}
	}

	return inverseMatrix, nil
}
