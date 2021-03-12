package algebra

func AdjointMatrix(matrix [][]float64) (adjMatrix [][]float64, err error) {
	size := len(matrix)

	adjMatrix = CreateMatrix(size, size)
	if size == 1 {
		adjMatrix[0][0] = 1
		return adjMatrix, nil
	}
	sign := 1
	var subMatrix = CreateMatrix(size, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			Cofactor(matrix, subMatrix, i, j, size)
			if (i+j)%2 == 0 {
				sign = 1
			} else {
				sign = -1
			}
			adjMatrix[j][i] = float64(sign) * (Determinant(subMatrix, size-1))
		}
	}

	return adjMatrix, err
}
