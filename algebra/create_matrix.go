package algebra

func CreateMatrix(row, col int) [][]float64 {
	matrix := make([][]float64, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]float64, col)
	}

	return matrix
}
