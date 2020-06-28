package main

import "fmt"

/**
map使用哈希表，可以比较相等，存储的值是无序的
除了slice,function,map的内建类型都可以作为key
struck类型不包含上述字段，也可以作为key
*/
func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println("m = ", m)

	m2 := make(map[string]int) // m2 == empty map
	fmt.Println("m2 = ", m2)

	var m3 map[string]int // m3 = nil
	fmt.Println("m3 = ", m3)

	fmt.Println("Traversing map")
	// map遍历时是无序的
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Geting Value")
	// 通过第二个值可以确定键是否存在
	courseName, ok := m["course"]
	fmt.Println("courseName =", courseName, "ok =", ok)

	// 错误的键，取到的是空值
	caurseName, ok := m["caurse"]
	fmt.Println("caurseName =", caurseName, "ok =", ok)

	fmt.Println("Deleting Value")
	name, ok := m["name"]
	fmt.Println(ok, name)

	delete(m, "name")

	name, ok = m["name"]
	fmt.Println(ok, name)
}
