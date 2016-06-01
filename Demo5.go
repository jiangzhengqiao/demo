package main

import (
	"fmt"
	"strconv"
)

// 计算机存储单位的枚举
// B是计算机最小单位,iota默认为0，B就左移0为，KB移动1*10位
const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("Demo5")
	fmt.Println(YB)
	a := 1
	a++
	fmt.Println(a)
	b := 1
	b--
	fmt.Println(b)

	// if
	c := 100
	if c, d := 1, 2; c > 0 && d > 1 {
		fmt.Println("ok")
		fmt.Println(c)
		fmt.Println(d)
	}
	fmt.Println(c)

	// for
	a1 := 0
	for {
		a1++
		fmt.Println("a1=" + strconv.Itoa(a1))
		if a1 >= 3 {
			break
		}
	}
	fmt.Println("循环1结束")

	a2 := 0
	for a2 < 3 {
		a2++
		fmt.Println("a2=" + strconv.Itoa(a2))
	}
	fmt.Println("循环2结束")

	for a3 := 0; a3 < 3; a3++ {
		fmt.Println("a3=" + strconv.Itoa(a3))
	}
	fmt.Println("循环3结束")

	// switch
	s := 1
	switch s {
	case 0:
		fmt.Println("s=" + strconv.Itoa(s))
	case 1:
		fmt.Println("s=" + strconv.Itoa(s))
	default:
		fmt.Println("switch mei you !")
	}

	switch {
	case s >= 0:
		fmt.Println("s=" + strconv.Itoa(s))
		fallthrough
	case s >= 1:
		fmt.Println("s=" + strconv.Itoa(s))
	default:
		fmt.Println("switch mei you !")
	}

	switch s1 := 1; {
	case s1 >= 0:
		fmt.Println("s1=" + strconv.Itoa(s))
		fallthrough
	case s1 >= 1:
		fmt.Println("s1=" + strconv.Itoa(s))
	default:
		fmt.Println("switch mei you !")
	}
	// 变量已经不存在了，因为实在switch中定义的
	// fmt.Println("s1=" + strconv.Itoa(s1))

	// 跳转
LABLE1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABLE1
			}
		}
	}
	fmt.Println("break 跳出了循环！")

	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LABLE2
			}
		}
	}
LABLE2:
	fmt.Println("goto 跳出了循环！")

	// continue 只可以用在有限的循环
LABLE3:
	for i := 0; i < 10; i++ {
		for {
			continue LABLE3
		}
	}
	fmt.Println("continue 跳出了循环！")
}
