package main

import (
	"fmt"
	"github.com/ajantonov/machine-learning/statistics"
)

func main() {
	fmt.Println("Hello Machine World!")
	values := []float64{1}

	result, _ := statistics.Mean(values)

	fmt.Println("Result : ", result)
}
