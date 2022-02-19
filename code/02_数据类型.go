package main

import (
	"fmt"
	"math"
)

/*
基本数据类型:
	bool                                           // 布尔
	string                                         // 字符串
	(u)int, (u)int8, (u)int16, (u)int32, (u)int64  // 有u代表无符号, 无u带表有符号
	uintptr                                        // 指针
	float32, float64                               // 浮点数
	complex64, complex128                          // 复数

注意: 数据类型转换必须使用强制类型转换。
*/

func main() {
	var (
		a bool      = true
		b string    = "string type"
		c int       = -20
		d uint      = 20
		e float32   = 3.14
		f complex64 = 3 + 2i
	)
	fmt.Println("基本数据类型:")
	fmt.Println(a, b, c, d, e, f) // true string type -20 20 3.14 (3+2i)

	var var1 int32 = 100
	var var2 int64
	var2 = int64(var1)
	fmt.Println("强制类型转换:")
	fmt.Println(var1, var2) // 100 100

	// 实例: 三角函数
	var aa, bb = 3, 4
	fmt.Println(math.Sqrt(float64(aa*aa + bb*bb)))  // 5
}

