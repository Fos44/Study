package main

import "fmt"

const n = 3

func sum(x []int64)float64  {
	var s float64 = 0
	for i := 0; i < len(x) - 1; i++{
		s = s + float64(x[i])
	}
	return s
}

func main() {
	var a [n][]int64
	var avg = [n]float64{}

	fmt.Println("Input data")

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &a[i][j])

		}
	}
	fmt.Println(avg)
	}

