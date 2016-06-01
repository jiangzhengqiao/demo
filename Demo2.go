package main

import (
	"fmt"
)

//常亮
const SHUJU = "你说啥"

// 全局变量
var status = 1

// 类型
type newType int

// 结构
type gopher struct{}

// 接口
type golang interface{}

func main() {
	fmt.Println("Demo2")
	fmt.Println(SHUJU)
}
