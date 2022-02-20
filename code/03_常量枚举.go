package main

import (
	"fmt"
)

/*
1. 定义常量
	常量定义之后不引用也可以通过编译。

2. 枚举
	const() 定义
	const([var] = iota) 定义，每一次const iota初始化为0

注意：其他语言习惯常量或枚举使用大写，在Go语言中首字母大写默认Public其他包可见。
*/

func main() {
	fmt.Println("1. 定义常量:")
	const pi = 3.14
	const os_version = "centos7.0"
	const a, b = 3, 4
	const (
		c = 5
		d = 6
	)
	fmt.Println(pi, os_version, a, b, c, d) // 3.14 centos7.0 3 4 5 6

	fmt.Println("2. 枚举:")
	const (
		langC      = "C"
		langJava   = "JAVA"
		langPython = "Python"
		langGo     = "GO"
	)

	const (
		n1 = iota
		n2
		_
		n4
	)
	fmt.Println(langC, langJava, langPython, langGo) // C JAVA Python GO
	fmt.Println(n1, n2, n4)                          // 0 1 3
}

