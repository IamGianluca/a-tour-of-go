package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var cnt int
	var z float64
    var tol float64
    for z = 1; cnt < 10; cnt += 1 {
		tol = (z*z - x) / (2*z)
        z -= tol
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
