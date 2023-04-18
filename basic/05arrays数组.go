package main

import "fmt"

/*
*
array是值类型
*/
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Printf("arr[%d]:{} %d\n", i, v)
	}
}

/*
*
通过传递数组指针改变元素
*/
func printArrayPointer(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Printf("arr[%d]:{} %d\n", i, v)
	}
}

func main() {
	var arr1 [5]int                  // 定义长度5的数组
	arr2 := [3]int{1, 3, 5}          // 冒号等于要指定初始值
	arr3 := [...]int{2, 4, 6, 8, 10} //可变长数组
	fmt.Println(arr1, arr2, arr3)

	// 二维数组
	var grid [4][5]int
	fmt.Println(grid)

	// 遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("arr[%d]:{} %d\n", i, arr3[i])
	}
	// 使用range关键字
	for i, v := range arr3 {
		fmt.Printf("arr[%d]:{} %d\n", i, v)
	}

	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("printArray(arr1)")
	printArray(arr3)

	fmt.Println("fmt.Println(arr1, arr3)")
	fmt.Println(arr1, arr3)

	fmt.Println("printArrayPointer(arr1)")
	printArrayPointer(&arr1)

	fmt.Println("printArrayPointer(arr1)")
	printArrayPointer(&arr3)

	fmt.Println("fmt.Println(arr1, arr3)")
	fmt.Println(arr1, arr3)
}
