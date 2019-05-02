package main

import "fmt"
import "math"

func f(x float64)float64  {
	return math.Pow(x,2) - 24 * x + 14
}

func main() {


	var f1 float64 = 1
	var f2 float64 = 1
	var f3 float64 = f1 + f2
	var alpha = 0.0005
	var eps = 0.00001
	var a float64 = -2
	var b float64 = 16

	for f3 < ((b - a) / (2 * eps)){
		f1 = f2
		f2 = f3
		f3 = f1 + f2
	}

	var x = a + (f1 / f3) * (b - a)
	var y = a + b - x

	if f(x) < f(y){
		b = y
		y = x
		x = a + b - y
	}else {
		a = x
		x = y
		y = a + b - y
	}

	var res float64

	for math.Abs(x - y) > alpha  {
		if f(x) < f(x + alpha) {
			res = (a + b + alpha) / 2
		}else {
			res = (a + b) / 2
		}

	}

fmt.Println("Точка: ", res)

}
