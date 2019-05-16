package main

import "fmt"

func main() {
	const n = 27
	var a = [n]float64{}
	var res = [n]int{}
	var r float64 = 6

	fmt.Println("Input data")

	for i := 0; i < n; i++{
		fmt.Scan(&a[i])
		res[i] = int(a[i] * r) + 1
	}

	fmt.Println(res)
}
