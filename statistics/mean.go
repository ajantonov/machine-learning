/*
 This package contains fundamental functions used for statistics learning.
*/
package statistics

import "errors"

/*
Mean
	The sum of all values divided by the number of values.

	Synonymous
		average
*/
func Mean(values []float64) (float64, error) {

	size := len(values)

	if 0 == size {
		return 0, errors.New("invalid input parameter")
	}

	var result = float64(0)

	for _, value := range values {
		result += value
	}

	result /= float64(size)

	return result, nil
}
