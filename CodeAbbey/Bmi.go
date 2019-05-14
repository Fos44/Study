package main

import (
	"fmt"
	"math"
)
/*
func bmi(x float64, y float64)string{
	var index = x / math.Pow(y, 2)
	var norma string
	if index < 18.5{
		norma = "under"
	}else if 18.5 <= index && index < 25 {
		norma = "normal"
	}else if 25 <= index && index < 30 {
		norma = "over"
	}else if 30 <= index {
		norma = "obese"
	}
	return norma
*/

func main() {

	const n = 3
	var a = [n]float64{80, 55, 49}
	var b = [n]float64{1.73, 1.58, 1.91}
	var index = [n]float64{}

	var norma = [n]string{}

	fmt.Println("Введите данные")

	for i := 0;i < n ;i++  {
		index[i] = a[i] / math.Pow(b[i], 2)
		if index[i] < 18.5{
			norma[i] = "under"
		}else if index[i] >= 18.5 && index[i] < 25 {
			norma[i] = "normal"
		}else if index[i] >= 25 && index[i] < 30 {
			norma[i] = "over"
		}else if  index[i] >= 30 {
			norma[i] = "obese"
		}
	}
	fmt.Println(norma)
}
