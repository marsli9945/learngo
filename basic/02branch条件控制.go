package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

/**
switch语句每个case后默认有break
*/
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"

	/**
	if条件中可以赋值
	if条件中的变量作用域只在if语句代码块中
	*/
	if contents, err := ioutil.ReadFile(filename); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		grade(1001),
	)
}
