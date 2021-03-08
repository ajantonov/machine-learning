package algebra

func AddMatrices(matrixOne [][]float64, matrixTwo [][]float64) [][]float64 {

	var result = make([][]float64, len(matrixOne))

	for i, row := range matrixOne {
		for j := range row {
			result[i] = append(result[i], matrixOne[i][j]+matrixTwo[i][j])
		}
	}

	return result
}
