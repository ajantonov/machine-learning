module github.com/ajantonov/machine-learning

go 1.14

replace (
	github.com/ajantonov/machine-learning/algebra => ./algebra
	github.com/ajantonov/machine-learning/statistics => ./statistics
)

require gonum.org/v1/gonum v0.9.0
