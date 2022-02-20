package main

import "fmt"

/*
1. 条件语句
	if...else if...else
	switch
*/

func main() {
	fmt.Println("1. if...else语句:")
	var num float32 = 12
	if num >= 0 && num <= 10 {
		fmt.Println("0 <= num <= 10")
	} else if num > 10 {
		fmt.Println("num > 10")
	} else {
		fmt.Println("num < 0")
	}
	fmt.Println()

	fmt.Println("2. switch语句:")
	switch {
	case num >= 0 && num <= 10:
		fmt.Println("0 <= num <= 10")
	case num > 10:
		fmt.Println("num > 10")
	default:
		fmt.Println("num < 0")
	}
	fmt.Println()
}

