package main

import "fmt"

func main() {
	const n = 23
	var a = [n]int{}
	var b = [n]int{}
	var res = [n]int{}
	var r = 6

	fmt.Println("Input data")

	for i := 0; i < n; i++{
		fmt.Scan(&a[i],&b[i])
		res[i] = (a[i] % r + 1) + (b[i] % r + 1)
	}
	fmt.Println(res)
}
