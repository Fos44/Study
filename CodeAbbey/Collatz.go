package main

import "fmt"

func main() {

	const n = 25
	var a =[n]int64{}
	var res = [n]int64{}

	fmt.Println("Input data")

	for i := 0; i < n; i++{
		fmt.Scan(&a[i])
		for a[i] != 1 {
			if a[i]%2 == 0 {
				res[i] += 1
				a[i] /= 2
			} else {
				a[i] = (3 * a[i]) + 1
				res[i] += 1
			}
		}
	}
	fmt.Println(res)
}
