package main

import "fmt"
import "math"

var x0 = []float64{1.5, 0.1}

func fu(x_ []float64) float64 {
	var ff =math.Pow(x_[0], 2) + 8 * math.Pow(x_[1], 2) - x_[0] * x_[1] + x_[0]
	return  ff
}

func grad1(x0[]float64) float64  {
	return  2 * x0[0] - x0[1] + 1
}

func grad2(x0[]float64)float64  {
	return 16 * x0[1] - x0[0]
}

func grad(x0 []float64)  []float64{
	var g = []float64{grad1(x0),grad2(x0)}
	return g
}

//новый вектор с помощью антиградиента
func newX (x0[]float64, t float64) []float64 {
	var nX = []float64 {x0[0] - t * grad (x0)[0] , x0[1] - t * grad (x0)[1]}
	return nX
}

func main() {
//	var x0 = []float64{1.5, 0.1}
	var eps1 = 0.1

	var coverage = false
	var st = 0

	//золотое сечение
	var a float64 = 10
	var b float64 = -10
	var rat = (3 - math.Sqrt(5)) / 2
	var x1  = a + rat * (b - a)
	var x2 = a + b - x1

	var xElder = []float64{0,0}
	var x = x0

	// Основной цикл
	for !coverage   {

		st++
		fmt.Println("Step:", st)

		if fu(newX ( x ,x1 ) ) < fu( newX( x, x2 ) ) {
				b = x2
				x2 = x1
				x1 = a + b - x2
				x = newX(x, x1)
			} else {
					a = x1
					x1 = x2
					x2 = a + b - x1
					x = newX(x, x2)
		}
		//Норма вектора
			var norma = math.Sqrt(math.Pow(xElder[0]- x[0],2) + math.Pow(xElder[1] - x[1],2))
		fmt.Println("Norma:",norma)
		xElder = x
		coverage = norma <= eps1
		fmt.Println("X:",x)

	}

}
