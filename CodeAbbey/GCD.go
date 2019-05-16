package main

import (
	"fmt"
)

func main() {

	const n = 23
	var a = [n]float64{}
	var b = [n]float64{}
	var aa = [n]float64{}
	var bb = [n]float64{}
	var res = [n][2]float64{}
	var gcd = [n]float64{}
	var lcm = [n]float64{}

	fmt.Println("Input data")

	for i := 0; i < n; i++{
		fmt.Scan(&a[i],&b[i])
		aa[i] = a[i]
		bb[i] = b[i]
		for a[i] != b[i]{
			if a[i] > b[i]{
				a[i] -= b[i]
			}else {
				b[i] -= a[i]
			}
		gcd[i] = a[i]
		}
		lcm[i] = (aa[i] * bb[i]) / gcd[i]
	res[i][0] = gcd[i]
	res[i][1] = lcm[i]
	}
	fmt.Println(res)
}
