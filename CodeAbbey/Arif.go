package main

import "fmt"

func main() {

	var a = [8] int{10, 24, 30, 7, 13, 25, 12, 19}
	var b = [8] int{18, 9, 16, 15, 10, 0, 8, 3}
	var c = [8] int{46, 97, 39, 19, 62, 56, 77, 79}

	var d [8] int

	for i := 0; i < len(a); i++{

		var s = a[i]

		for j := 0; j < c[i]; j++{
			s = s + b[i]
		}

		d[i] = ((c[i]) * (a[i] + s)) / 2
	}

	fmt.Println("Результат", d)

	
}
