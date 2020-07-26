package tree

/**
要改变内容必须使用指针接收者
结构过大也考虑使用指针接收者
一致性：如果有指针接收者，最好都是指针接收者

值接收者 是go语言特有
值/指针接收者均可接收值/指针接收者
*/
func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
