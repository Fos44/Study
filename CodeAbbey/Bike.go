package main

import "fmt"

func main() {

	const n = 1

	var a = [n]float64{}
	var b = [n]float64{}
	var dis = [n]float64{}
	var res = [n]float64{}

	fmt.Println("Input data")

	for i := 0; i < n; i++{
		fmt.Scan(&dis[i], &a[i], &b[i])
		res[i] = dis[i] / float64(a[i] + b[i])
		res[i] *= a[i]
	}
	fmt.Println(res)
}
