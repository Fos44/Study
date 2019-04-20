package main

import "fmt"
import "math"

var eps = 0.5
var a float64 = -2
var b float64 = 16
var s  = 0


var ratio = (3 - math.Sqrt(5))/2

var x = a + ratio * (b - a)
var y = a + b - x

var flag = true

func f( x float64) float64  {
	return math.Pow(x, 2) - (24 * x) - 14
}

func main() {
	for math.Abs(b - a) > 2 * eps  {
		if f(x) < f(y) {
				b = y
				y = x
				x = a + b - y
		} else {
				a = x
				x = y
				y = a + b - x
				}
		s = s + 1
		}
	s = s + 1
 	var result = (a + b) / 2
	fmt.Println ("Количество шагов:", s)
	fmt.Println ("Результат:", result)
	}


