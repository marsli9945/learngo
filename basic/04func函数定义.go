package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (num int, err error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

/**
带余除法
13 / 3 = 4 余 1
函数可以多返回值
返回值定义名称后
内部可以直接赋值后return
外部编辑器可以通过快捷键自动赋值
*/
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

/**
函数可以作为参数传入
在内部调用
*/
func apply(op func(int, int) int, a, b int) int {

	pointer := reflect.ValueOf(op).Pointer()  // 获取op的指针
	name := runtime.FuncForPC(pointer).Name() // 通过指针获取op的名字

	fmt.Printf("Calling function %s with args(%d, %d)\n", name, a, b)

	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

/**
通过传递指针实现引用传递
*/
func swap(a, b *int) {
	*a, *b = *b, *a
}

/**
返回值类型写在最后面
可返回多个值
函数可以作为参数和返回值
没有 默认参数、可选参数
*/
func main() {
	num, err := eval(3, 4, "+")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(num)
	}
	//fmt.Println(
	//	eval(1, 2, "/"),
	//	eval(12, 2, "/"),
	//)

	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))
	// 传入匿名函数
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
