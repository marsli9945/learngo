package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

/**
go 没有构造函数
初始化使用自定义工厂函数，返回构造好对结构体即可
注意返回了局部变量地址，这在go中是合法对
*/
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

/**
显式定义和命名方法接收者
只有使用指针才可以改变结构体内容
nil指针也能调用方法
*/
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}

	node.Value = value
}
