package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64 = 3
	var y float64 = 4
	hyp := hypot(x, y)

	fmt.Printf("hypot: %T\n", hyp)
	fmt.Printf("x: %f; y: %f; hypotenuse: %f", x, y, hyp)
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
