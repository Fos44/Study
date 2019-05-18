package main

import "fmt"

func main() {
	const n = 289
	const t = 20
	var a [n]int64
	var res[t] int64

	fmt.Println("Input data")

	for i :=0 ;i < n; i++ {
		fmt.Scan(&a[i])
			switch a[i] {
			case 1:
				res[0] = res[0] + 1
			case 2:
				res[1] = res[1] + 1
			case 3:
				res[2] = res[2] + 1
			case 4:
				res[3] = res[3] + 1
			case 5:
				res[4] = res[4] + 1
			case 6:
				res[5] = res[5] + 1
			case 7:
				res[6] = res[6] + 1
			case 8:
				res[7] = res[7] + 1
			case 9:
				res[8] = res[8] + 1
			case 10:
				res[9] = res[9] + 1
			case 11:
				res[10] = res[10] + 1
			case 12:
				res[11] = res[11] + 1
			case 13:
				res[12] = res[12] + 1
			case 14:
				res[13] = res[13] + 1
			case 15:
				res[14] = res[14] + 1
			case 16:
				res[15] = res[15] + 1
			case 17:
				res[16] = res[16] + 1
			case 18:
				res[17] = res[17] + 1
			case 19:
				res[18] = res[18] + 1
			case 20:
				res[19] = res[19] + 1
			}

	}
	fmt.Println(res)
}
