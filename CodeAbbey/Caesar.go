package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const n = 2
	var s string
	fmt.Println("Input data")

	for i := 0; i < n; i++{
		myscanner := bufio.NewScanner(os.Stdin)
		myscanner.Scan()
		s = myscanner.Text()

		strings.Replace(s, "D", "A", 1)
		strings.Replace(s, "E", "B", 1)
		strings.Replace(s, "F", "C", 1)
		strings.Replace(s, "G", "D", 1)
		strings.Replace(s, "H", "E", 1)
		strings.Replace(s, "I", "F", 1)
		strings.Replace(s, "J", "G", 1)
		strings.Replace(s, "K", "H", 1)
		strings.Replace(s, "L", "I", 1)
		strings.Replace(s, "M", "J", 1)
		strings.Replace(s, "N", "K", 1)
		strings.Replace(s, "O", "L", 1)
		strings.Replace(s, "P", "M", 1)
		strings.Replace(s, "Q", "N", 1)
		strings.Replace(s, "R", "O", 1)
		strings.Replace(s, "S", "P", 1)
		strings.Replace(s, "T", "Q", 1)
		strings.Replace(s, "U", "R", 1)
		strings.Replace(s, "V", "S", 1)
		strings.Replace(s, "W", "T", 1)
		strings.Replace(s, "X", "U", 1)
		strings.Replace(s, "Y", "V", 1)
		strings.Replace(s, "Z", "W", 1)
		strings.Replace(s, "A", "X", 1)
		strings.Replace(s, "B", "Y", 1)
		strings.Replace(s, "C", "Z", 1)

	}

	fmt.Println(s)
}