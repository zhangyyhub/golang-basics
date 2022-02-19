package main

import "fmt"

/*
变量定义:
	1. 单个变量定义
		方法0: 简洁语法
		方法1: var [变量名] [类型] = [value]
		方法2: var [变量名] = [value]
		方法3: var [变量名] [类型]

		注意: 其中简洁语法不能在全局中定义，只能在函数体内部定义。
	2. 多个变量定义
*/

func main() {
	// 1. 单个变量定义
	name0 := "xiaoming"         // 方法0: 简洁语法
	var name1 string = "yannic" // 方法1: var [变量名] [类型] = [value]
	var name2 = "jack"          // 方法2: var [变量名] = [value]
	var name3 string            // 方法3: var [变量名] [类型]

	fmt.Println("1. 单个变量定义")
	fmt.Printf("%s %s %s %q\n", name0, name1, name2, name3) //xiaoming yannic jack ""

	// 2. 多个变量定义
	var name4, name5 string = "yang", "zhang"
	var name6, name7 = "wang", "zhao"
	var name8, name9 string

	var (
		age1 int = 22
		age2     = 18
		age3 int
	)

	fmt.Println("2. 多个变量定义")
	fmt.Printf("%s %s %s %s %q %q\n", name4, name5, name6, name7, name8, name9) // yang zhang wang zhao "" ""
	fmt.Printf("%d, %d, %d\n", age1, age2, age3)                                // 22 18 0
}
