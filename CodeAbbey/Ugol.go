package main

import (
	"fmt"
	"math"
)

func main() {

	const n = 26
	var a = [n]float64{}
	var b = [n]float64{}
	var c = [n]float64{}

	var res = [n]string{}

	fmt.Println("Input data")

	for i := 0; i < n ; i++{
		fmt.Scan(&a[i],&b[i],&c[i])
		if math.Pow(c[i], 2) == math.Pow(a[i],2) + math.Pow(b[i],2){
			res[i] = "R"
		}else if math.Pow(c[i],2) < math.Pow(a[i],2) + math.Pow(b[i],2){
			res[i] = "A"
		}else if  math.Pow(c[i],2) > math.Pow(a[i],2) + math.Pow(b[i],2){
			res[i] = "O"
		}
	}
	fmt.Println(res)
	
}
