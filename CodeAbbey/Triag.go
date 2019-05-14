package main

import "fmt"

func main() {

	const n = 26

	var a int64
	var b int64
	var c int64

	var v = [n]int64{}

	fmt.Println("Введите точки")

	for i := 0; i < n; i++{
		fmt.Scanf("%d%d%d%d", &a, &b, &c)
		if (a < b + c) && (b < a + c) && ( c < a + b){
			v[i] = 1
		}else{
			v[i] = 0
		}

	}

	fmt.Println(v)
}
