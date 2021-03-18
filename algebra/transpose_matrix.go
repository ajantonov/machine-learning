package algebra

func TransposeMatrix(matrix [][]float64) (transposedMatrix [][]float64, err error) {

	rows := len(matrix)
	columns := len(matrix[0])

	transposedMatrix = CreateMatrix(columns, rows)

	for i, r := range matrix {
		for j, value := range r {
			transposedMatrix[j][i] = value
		}
	}

	return transposedMatrix, err
}
