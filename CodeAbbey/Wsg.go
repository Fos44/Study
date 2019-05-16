package main

import (
	"fmt"
	"math"
)

func main() {

	const n  = 1
	var a  int64
	var pr = [n]int64{}
	var res = [n]int64{}
	var s int64 = 0
	var f int64 = 0
	var prom = [n]int64{}

	fmt.Println("Введите числа")

	for i := 0; i < n; i++{
		fmt.Scan(&a)
		pr[i] = a
		prom[i] = a
		for prom[i] != 0{
			f++
			prom[i] /= 10
		}
		fmt.Println(f)
		for pr[i] != 0{
			s = s + 1
			res[i] += (pr[i] % int64(math.Pow(float64(10),float64(f - 2)))) * s
			pr[i] /= 10
			f = f - 1

		}

	}
	fmt.Println(res)
}
