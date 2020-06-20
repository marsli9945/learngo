package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
十进制转二进制
for循环省略起始条件
*/
func convertToBin(n int) string {

	if n == 0 {
		return "0"
	}

	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

/**
逐行读取文件
for循环省略起始条件和递增条件
*/
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/*
for循环省略起始条件、判断条件、递增条件
此时相当与常驻内存(死循环)
*/
func forever() {
	for {
		fmt.Println("forever")
	}
}

/*
for,if后面的条件没有括号
if条件里也可以定义变量
没有while
switch不需要break，也可以直接switch多个条件
*/
func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(0),
	)

	printFile("abc.txt")
	forever()
}
