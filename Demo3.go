package main

import (
	"fmt"
	"math"
)

const (
	SHUJU = "你说啥"
	WDZJ  = "再说一遍"
)

var (
	A = "A"
	B = "B"
)

type (
	文本       string
	ByteSize int64
)

func main() {
	fmt.Println("Demo3")
	fmt.Println(SHUJU)
	fmt.Println(WDZJ)
	fmt.Println(A)
	fmt.Println(B)

	// 类型零值（string和自定义类型冲突）
	var a int
	// var b string
	var c bool
	var d byte
	var e int32
	var f float32
	var g [1]int
	var h [1]bool
	var i [1]byte

	fmt.Println(a)
	// fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	// 检查是否超出范围
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MinInt8)

	// 自定义类型
	var b 文本 = "中文类型名"
	fmt.Println(b)

	// 多个变量赋值,冒号是代替var关键字的，同时出现会出错，使用_可以忽略赋值
	var a1, b1, c1, d1 int = 1, 2, 3, 4
	// var a1, b1, c1, d1 = 1, 2, 3, 4
	// a1, b1, c1, d1 := 1, 2, 3, 4
	fmt.Println("a1=" + string(a1))
	fmt.Println("b1=" + string(b1))
	fmt.Println("c1=" + string(c1))
	fmt.Println("d1=" + string(d1))

	// 变量类型转换
	var j1 float32 = 100.9
	fmt.Println(j1)
	j2 := int(j1)
	fmt.Println(j2)
}
