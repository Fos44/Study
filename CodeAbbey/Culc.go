package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var s [7]string
	var a [7]int
	var num int

	fmt.Println("input data")

	for i := 0; i > 0; i++{
		myscanner := bufio.NewScanner(os.Stdin)
		myscanner.Scan()
		s[i] = myscanner.Text()
		fmt.Scan(&num, &a[i])

		if s[i] == "+"{
			num += a[i]
		}else if s[i] == "*"{
			num *= a[i]
		}else {
			num %= a[i]
			break
		}
	}
	fmt.Println(num)
}
