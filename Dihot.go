package main

import "fmt"
import "math"

func f(x float64) float64  {
	return math.Pow(x,2) - 24 * x + 14
}

func main() {

	var a float64 = -2
	var b float64 = 16
	var e float64 = 0.5
	var d float64 = 0.2
	var z float64
	var y float64

	for math.Abs(b - a) > (2 * e) {

		var center = (a + b) / 2
		y = center - d
		z = center + d

		if f(y) > f(z){
			a = y
		}else {
			b = z
		}
	}
	fmt.Println("Точка: ", (b + a) / 2)
}
