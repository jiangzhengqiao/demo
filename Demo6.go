package main

import (
	"fmt"
	// "strconv"
)

func main() {
	fmt.Println("Demo6")

	// 数据长度不等，不能直接用等号赋值，如果长度相等，就可以用等号
	var a [2]int
	var b [2]int
	b = a
	fmt.Println(b)

	// 默认是0赋值
	a1 := [2]int{1}
	fmt.Println(a1)

	b1 := [20]int{19: 1, 5: 1, 10: 1}
	fmt.Println(b1)

	c1 := [...]int{1, 2, 3}
	fmt.Println(c1)

	d1 := [...]int{0: 1, 1: 10, 2: 100, 100: 1}
	fmt.Println(d1)

	// 指向数组的指针
	var p *[2]int = &a1
	fmt.Println(p)

	e1 := [2]int{1, 2}
	f1 := [2]int{1, 3}
	fmt.Println(e1 == f1)

	// 数组赋值
	p1 := new([10]int)
	p1[1] = 222
	fmt.Println(p1)

	p2 := [10]int{}
	p2[1] = 2
	fmt.Println(p2)

	// 多纬数组
	c3 := [...][3]int{
		{1, 1, 1},
		{2, 2, 2},
	}
	fmt.Println(c3)

	// 冒泡排序
	mp := [5]int{61, 23, 57, 82, 93}
	fmt.Println(mp)

	num := len(mp)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if mp[i] < mp[j] {
				temp := mp[i]
				mp[i] = mp[j]
				mp[j] = temp
			}
		}
	}
	fmt.Println(mp)
}
