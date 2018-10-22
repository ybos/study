package imooc_package_node

import "fmt"

type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func (node *TreeNode) Print() {
	fmt.Println(node.Value)
}