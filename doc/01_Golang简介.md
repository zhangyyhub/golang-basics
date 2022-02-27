# Golang 简介

Go（又称 Golang）是 Google 的 Robert Griesemer，Rob Pike 及 Ken Thompson 开发的一种静态强类型、编译型语言。Go 语言语法与 C 相近，但功能上有：内存安全，GC（垃圾回收），结构形态及 CSP-style 并发计算。

官网：https://golang.org/

# 1、Go 语言特点

1. 背靠大厂（Google），可靠
2. 天生支持并发（最显著特点）
3. 语法简单，容易上手
4. 内置 runtime，支持垃圾回收
5. 可直接编译成机器码，不依赖其他库
6. 丰富的标准库
7. 跨平台编译
8. 部署维护成本低

# 2、Go 语言应用邻域

1. 服务器编程
2. 开发云平台
3. 区块链
4. 分布式系统
5. 网络编程

# 3、使用 Go 语言的公司

1. Google

	k8s

2. Facebook

	facebookgo

3. 腾讯

	蓝鲸平台、容器技术

4. 百度

	运维项目 BFE

5. 京东

	消息推送系统、云存储、京东商城

6. 小米

	运维监控系统、小米互娱、小米商城、小米视频、小米生态链

7. 360

	日志搜索系统 Poseidon

# 4、Go 语言代码风格

go 应用使用包和模块来组织代码，包对应到文件系统就是文件夹，模块就是 .go 的 go 源文件。一个包中会有多个模块，或者多个子包。

# 5、Go 项目管理工具

早期的 go 项目使用 gopath 来管理项目，不方便而且容易出错。从 golang 1.11 开始使用 gomod 管理项目，当然还有第三方模块例如 govendor 等。

1. 创建项目路径

	一般我们定义项目路径为：`[domain]/[project-name]`，例如：`zhangyyhub.com/pro-basics`

2. 初始化模块，生成 go.mod 文件

~~~bash
➜ cd zhangyyhub.com/pro-basics
➜ go mod init zhangyyhub.com/pro-basics
➜ cat go.mod
module zhangyyhub.com/pro-basics

go 1.17
~~~

3. 创建 zhangyyhub.com/pro-basics/demo.go 文件

~~~go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, world.")
}
~~~

4. 运行 go 程序

**方法一：** 编译并运行 Go 程序

~~~bash
➜ cd zhangyyhub.com/pro-basics
➜ go run demo.go
hello, world.
~~~

**方法二：** 编译并安装包和依赖项

~~~bash
➜ cd zhangyyhub.com/pro-basics
➜ go install zhangyyhub.com/pro-basics
➜ ll /usr/local/go-packages/bin/pro-basics
-rwxr-xr-x. 1 root root 1.7M Feb 26 12:45 /usr/local/go-packages/bin/pro-basics
➜ /usr/local/go-packages/bin/pro-basics
hello, world.
~~~

5. 子包的应用

**创建子包：**

~~~bash
➜ mkdir zhangyyhub.com/pro-basics/morestrings
➜ cat zhangyyhub.com/pro-basics/morestrings/morestrings.go
package morestrings

//子包
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
~~~

**应用子包：**
~~~bash
➜ cat zhangyyhub.com/pro-basics/demo.go
package main

import (
	"fmt"

	"zhangyyhub.com/pro-absics/morestrings"
)

func main() {
	fmt.Println("hello, world.")

	//调用子包
	s := morestrings.ReverseRunes("12345")
	fmt.Printf("s: %v\n", s)
}

➜ go run zhangyyhub.com/pro-basics/demo.go
hello, world.
s: 54321
~~~

6. 导入网络模块

查找地址：https://pkg.go.dev/

**安装：** `go get -u github.com/google/go-cmp/cmp`

**应用：**

~~~bash
➜ cd zhangyyhub.com/pro-basics

# 查看 mod 变量
➜ go mod tidy
module zhangyyhub.com/pro-absics

go 1.17

require github.com/google/go-cmp v0.5.7

# 应用网络模块
➜ cat demo.go
package main

import (
	"fmt"

	"zhangyyhub.com/pro-absics/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("hello, world.")

	//调用子包
	s := morestrings.ReverseRunes("12345")
	fmt.Printf("s: %v\n", s)

	//使用第三方包
	fmt.Println(cmp.Diff("Hello, World.", "Hello, Go."))
}

# 编译并运行
➜ go run demo.go 
hello, world.
s: 54321
  string(
-       "Hello, World.",
+       "Hello, Go.",
  )
~~~