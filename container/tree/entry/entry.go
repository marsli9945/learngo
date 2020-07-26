package main

import (
	"fmt"
	"github.com/marsli9945/learngo/container/tree"
)

/**
扩展类型有两种方式
定义别名
使用组合
*/
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

/**
结构体与方法都是首字母大写为公开，小写为私有
为结构体定义的方法必须放在同一包内
可以是不同文件
*/
func main() {
	root := tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()

	fmt.Println("myNode.......")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}
