package main

import "fmt"

/**
[]int 表示参数是slice
slice可以视为是数组的指针
可以改变内部的值
*/
func updateSlice(s []int) {
	s[0] = 100
}

/**
通过传递slice改变array的元素
*/
func printSlice(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Printf("arr[%d]:{} %d\n", i, v)
	}
}

/**
slice是array的view，可以通过slice改变array的元素值
slice可以向后扩展，但不可以向前扩展
s[i]不可以超过len(s),向后扩展不可超过底层array的长度cap(s)
向slice添加元素时如果超出cap，系统会重新分配更大的底层array
由于值传递的关系，必须接受append的返回值
s = append(s, val)
*/
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7} // 先定义数组

	// 切片是数组的视图
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1, s2, arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s1, s2, arr)

	var arr1 [5]int                  // 定义长度5的数组
	arr3 := [...]int{2, 4, 6, 8, 10} //可变长数组

	fmt.Println("printSlice(arr1)")
	printSlice(arr1[:])

	fmt.Println("printSlice(arr1)")
	printSlice(arr3[:])

	fmt.Println("fmt.Println(arr1, arr3)")
	fmt.Println(arr1, arr3)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending Slice")
	fmt.Println("arr", arr)
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1,
		len(s1),
		cap(s1),
	)
	s2 = s1[3:5]
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2,
		len(s2),
		cap(s2),
	)
	//fmt.Println(s1[4]) // 无法直接取s1[4]

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	// s4,s5 超出了arr的最大长度，不再是arr的view
	fmt.Println("arr = ", arr)
}
