package main

import "fmt"

func printSlice2(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Create Slice")
	var s []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice2(s)
		s = append(s, i*2+1)
	}

	// 使用冒号创建，指定初始值
	s1 := []int{2, 4, 6, 8}
	printSlice2(s1)

	// 使用make函数，创建长度为16的空slice
	s2 := make([]int, 16)
	printSlice2(s2)

	// 使用make函数创建时还可以指定cap值
	s3 := make([]int, 10, 32)
	printSlice2(s3)

	fmt.Println("Copy Slice")
	copy(s2, s1)
	printSlice2(s2)

	fmt.Println("Deleting elements from slice")
	// 删除s2[3],把4-15 append 到 0-2 再赋值给s2
	s2 = append(s2[:3], s2[4:]...)
	printSlice2(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice2(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice2(s2)
}
