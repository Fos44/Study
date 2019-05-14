package main

import "fmt"

func D(x, y, z int64)int64  {
	dg := (x * y) + z
	return dg
}

func main() {

	const n  = 12
	var a  int64
	var b int64
	var c int64
	var pr = [n]int64{}
	var res = [n]int64{}

	fmt.Println("Введите числа")

	for i := 0; i < n; i++{
		fmt.Scanf("%d%d%d%d", &a, &b, &c)
		pr[i] = D(a, b, c)
		for pr[i] != 0{
			res[i] += pr[i] % 10
			pr[i] /= 10
		}

	}
	fmt.Println(res)
}
