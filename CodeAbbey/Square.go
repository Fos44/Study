package main

import (
	"fmt"
	"math"
)

func main() {

	const n = 3
	var a = [n]float64{}
	var b = [n]float64{}
	var res = [n]float64{}
	var j float64
	var r float64 = 1
	var d float64

	fmt.Println("Input data")

	for i := 0; i < n; i++{
		fmt.Scan(&a[i],&b[i])
		for j = 0; j < b[i]; j++{
			res[i] = (r + a[i]/r) / 2
		}
		res[i] = math.Floor(res[i])
	}
	fmt.Println(res)
	
}
