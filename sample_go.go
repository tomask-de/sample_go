package sample_go

import "github.com/tomask-de/sample_go/newmath"

var result float64

func Setup() {
	result = newmath.Sqrt(4)
}

func GetResult() float64 {
	return result
}
