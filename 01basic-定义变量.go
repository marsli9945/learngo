package main

import "fmt"

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

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitValue()
	variableTypeDeuction()
}
