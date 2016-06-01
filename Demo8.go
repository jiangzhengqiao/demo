package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Demo8")
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[:]
	fmt.Println("s2 =", s2)

	//创建map
	var m map[int]string
	m = map[int]string{}
	m = make(map[int]string)
	fmt.Println("m =", m)

	m1 := make(map[int]string)
	m1[1] = "OK"
	m1[2] = "ERROR"
	delete(m1, 2)
	fmt.Println("m1 =", m1)

	// map赋值
	m2 := make(map[int]map[int]string)
	m2[1] = make(map[int]string)
	m2[1][1] = "OK"
	a := m2[1][1]
	fmt.Println("m2 =", m2)
	fmt.Println("a =", a)

	// 多返回值
	a1, ok := m2[2][1]
	if !ok {
		m2[2] = make(map[int]string)
		m2[2][1] = "GOOD"
	}
	a1, ok = m2[2][1]
	fmt.Println(a1, ok)

	// 批量初始化map
	sm := make([]map[int]string, 5)
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)

	// map 排序
	m3 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	// m3[1] = "a"
	// m3[2] = "b"
	// m3[3] = "c"
	// m3[4] = "d"
	// m3[5] = "e"
	fmt.Println("m3 =", m3)
	s3 := make([]int, len(m3))
	i := 0
	for k, _ := range m3 {
		s3[i] = k
		i++
	}
	fmt.Println("s3 =", s3)

	// KEY排序
	sort.Ints(s3)
	fmt.Println(s3)
}
