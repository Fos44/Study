package main

import "fmt"

func main() {
	const n = 24
	var a, b, c int64
	var m = [n]int64{}

	fmt.Println("Введите числа")

	for i := 0; i < n; i++{
		fmt.Scanf("%d%d%d%d", &a, &b, &c)
		if (a < b || a < c) && (a > c || a > b){
			m[i] = a
		}
		if (b < a || b < c) && (b > c || b > a){
			m[i] = b
		}
		if (c < a || c < b) && (c > b || c > a){
			m[i] = c
		}
	}

	fmt.Println(m)

}
