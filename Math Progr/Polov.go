package main

import "fmt"
import "math"

func f(x float64) float64 {
	return math.Pow(x,2) - 24 * x + 14
}

func main() {

	var i float64 = 1
	var a float64 = -2
	var b float64 = 16
	x := (a + b) / 2

	for b-a > i {
		y := a + math.Abs(b - a) / 4
		z := b - math.Abs(b - a) / 4
		if f(y) < f(z) {
			b = x
			x = y
		}else {
			if f(z) < f(x){
				a = x
				x = z
			}else {
				a = y
				b = z
			}
		}
	}
	fmt.Println("Точка: ", x)
}
