package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
函数外部定义变量
不能使用冒号
作用域为包内
*/
var (
	aa = 3
	ss = "kkk"
	bb = true
)

/**
变量定义
初始值打印
*/
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

/**
变量初值定义
同一行可定义多个变量
*/
func variableInitValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

/**
不指定定义的变量类型
由编译器自行推断
*/
func variableTypeDeuction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

/**
省略var关键字,使用冒号定义
只能在方法内使用
*/
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

/**
复数使用
验证欧拉公式
*/
func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1,
		//cmplx.Pow(math.E, 1i * math.Pi) + 1
	)
}

/**
类型转换
只有显示的类型转换
没有隐式的
*/
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

/**
定义常量
*/
func consts() {
	const (
		filename string = "abc.txt"
		a, b            = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

/**
常量枚举类型
iota表示自增
*/
func enums() {
	const (
		php = iota
		_
		java
		cpp
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(php, javascript, java, cpp)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

/**
变量类型写在变量名之后
编译器可推测变量类型
没有char，只有rune
原生支持复数类型
*/
func main() {
	fmt.Println("hello world")
	variableShorter()
	variableZeroValue()
	variableInitValue()
	variableTypeDeuction()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
	consts()
	enums()
}
