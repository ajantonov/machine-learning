package algebra

func NegativeMatrix(matrix [][]float64) [][]float64 {

	for i, row := range matrix {
		for j := range row {
			matrix[i][j] = -matrix[i][j]
		}
	}

	return matrix
}
