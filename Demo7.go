package main

import (
	"fmt"
)

// Slice
func main() {
	fmt.Println("Demo7")

	var s1 []int
	fmt.Println(s1)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(a)

	// 包含起止索引，不包含结束索引
	s2 := a[0:len(a)]
	fmt.Println(s2)

	s3 := a[0:]
	fmt.Println(s3)

	s4 := a[:len(a)]
	fmt.Println(s4)

	// 初始化slice，int类型，长度10，容量10,容量会翻倍扩容，在内存中分配一块连续的内存块，分配内存效率较低
	// 容量默认和长度一致
	s5 := make([]int, 3, 10)
	fmt.Println(s5, len(s5), cap(s5))

	b := []byte{'a', 'b',
		'c', 'd',
		'e', 'f',
		'g',
		'h',
		'i',
		'j',
		'k'}
	s6 := b[2:5]
	fmt.Println(string(s6))

	s7 := s6[1:3]
	fmt.Println(string(s7))

	s8 := make([]int, 3, 6)
	fmt.Printf("%p\n", s8)
	s8 = append(s8, 1, 2, 3)
	fmt.Printf("%v %p\n", s8, s8)
	s8 = append(s8, 4, 5, 6)
	fmt.Printf("%v %p\n", s8, s8)

	s9 := []int{1, 2, 3, 4, 5, 6}
	s10 := []int{7, 8, 9}
	copy(s9, s10)
	fmt.Println(s9)
}

func test() {
	fmt.Println("test")
}
