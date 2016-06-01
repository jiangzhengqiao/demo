package main

import (
	"fmt"
	"strconv"
)

// 枚举，常量大写，内部常量开头加_
const (
	a1 = iota
	b1
	c1
	d1
)

func main() {
	fmt.Println("Demo4")
	// 类型转换
	var a int = 65
	fmt.Println(a)
	b := strconv.Itoa(a)
	fmt.Println("b=" + b)
	a, _ = strconv.Atoi(b)
	fmt.Println(a)

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)
}
