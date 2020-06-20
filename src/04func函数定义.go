package main

import "fmt"

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		q, _ := div(a, b)
		return q
	default:
		panic("unsupported operation: " + op)
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

func main() {
	fmt.Println(
		eval(1, 2, "/"),
		eval(12, 2, "/"),
	)

	q, r := div(13, 3)
	fmt.Println(q, r)
}
