package main

import "fmt"

/*
1. 算数运算符
	+, -, *, /, %
	注意：整数的除法结果还是整数
2. 关系运算符
	==, !=, >, <, >=, <=
3. 逻辑运算符
	&&, ||, !
	注意：&&同真则真、||有真则真
4. 位运算符
	&, |, ^, <<, >>
5. 赋值运算符
	=, +=, -=, *=, /=, %=, &=, |=, <<=, >>=
*/

func main() {
	fmt.Println("1. 算数运算符:")
	fmt.Println(5/2, 5.0/2) // 2 2.5

	fmt.Println()
	fmt.Println("2. 关系运算符:")
	fmt.Println(5+3 == 7) // false

	fmt.Println()
	fmt.Println("3. 逻辑运算符:")
	fmt.Println(7 > 4 && 5 > 3, 7 > 4 || 5 > 9) // true, true

	fmt.Println()
	fmt.Println("4. 位运算符:")
	var a, b int = 3, 4
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b) // 4 3

	fmt.Println()
	fmt.Println("5. 赋值运算符:")
	var x int = 10
	x += 1
	fmt.Println(x) // 11
}

