package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func main() {
	const n = 19
	var s =  [n]string{}
	var res = [n]int{}
	fmt.Println("Input data")

	for i := 0; i < n; i++{
		myscanner := bufio.NewScanner(os.Stdin)
		myscanner.Scan()
		s[i] = myscanner.Text()

		strings.Replace(s[i], " ", "", -1)
		res[i] += strings.Count(s[i],"a" )
		res[i] += strings.Count(s[i],"o")
		res[i] += strings.Count(s[i],"u")
		res[i] += strings.Count(s[i],"i")
		res[i] += strings.Count(s[i],"e")
		res[i] += strings.Count(s[i],"y")

	}
	fmt.Println(res)
	
}
