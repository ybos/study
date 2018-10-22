package main

import "imooc_package"

type MyTreeNode struct {
	node *imooc_package_node.TreeNode
}

func (node *MyTreeNode) PostOrder() {
	if node == nil || node.node == nil {
		return
	}

	left := MyTreeNode{node.node.Left}
	right := MyTreeNode{node.node.Right}

	left.PostOrder()
	right.PostOrder()

	node.node.Print()
}

func main() {
	node := imooc_package_node.CreateNode(1)

	node.Left = imooc_package_node.CreateNode(2)
	node.Left.Left = imooc_package_node.CreateNode(3)

	node.Right = imooc_package_node.CreateNode(4)
	node.Right.Right = imooc_package_node.CreateNode(5)

	myNode := MyTreeNode{node}

	myNode.PostOrder()
}
