package algebra

func Determinant(matrix [][]float64, size int) float64 {
	var determinant = float64(0)

	if size == 1 {
		return matrix[0][0]
	}

	subMatrix := CreateMatrix(size, size)
	sign := 1

	for f := 0; f < size; f++ {
		Cofactor(matrix, subMatrix, 0, f, size)
		determinant += float64(sign) * matrix[0][f] * Determinant(subMatrix, size-1)
		sign = -sign
	}

	return determinant
}

func Cofactor(matrix [][]float64, cofactorMatrix [][]float64, rowIndex int, columnIndex int, size int) {
	var i, j = 0, 0

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if row != rowIndex && col != columnIndex {
				cofactorMatrix[i][j] = matrix[row][col]
				j++
				if j == size-1 {
					j = 0
					i++
				}
			}
		}
	}
}
