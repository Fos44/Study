package main

import "fmt"

func main() {
	const n = 1
	var day1[n] int
	var hour1[n] int
	var min1[n]int
	var sec1[n] int
	var day2[n] int
	var hour2[n] int
	var min2[n] int
	var sec2[n] int
	for i:=0; i<n; i++{
		fmt.Scan(&day1[i],&hour1[i],&min1[i],&sec1[i],&day2[i],&hour2[i],&min2[i],&sec2[i])
		var day int
		var hour int
		var minute int
		var second int
		day = day2[i]-day1[i]
		hour = ((hour2[i]*60)-(hour1[i]*60))/60
		minute = ((min2[i]*3600)-(min1[i]*3600))/3600
		second = ((sec2[i]*216000)-(sec1[i]*216000))/216000
		if second<0{
			second = second + 60
			minute = minute - 1
		}
		if minute<0{
			minute = minute + 60
			hour = hour - 1
		}
		if hour<0{
			hour = hour+24
			day = day - 1
		}
		fmt.Println("(",day,hour,minute,second,")");
	}
}