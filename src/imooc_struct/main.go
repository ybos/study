package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node treeNode) print() {
	fmt.Print(node.value)
}

// 默认是传值的方式来传递对象
// 值接收者
func (node treeNode) setValue(value int) {
	node.value = value
}

// 可以通过修改成指针对象来实现对对象的修改，同时指针和值的使用相同，无需修改
// 指针接收者
func (node *treeNode) setValue2(value int) {
	node.value = value
}

func main() {
	var root treeNode

	root1 := treeNode {value: 3}
	root1.left = &treeNode {}
	root1.right = &treeNode {5, nil, nil}
	root1.right.left = new(treeNode)
	root1.left.right = createNode(2)

	nodes := []treeNode {
		{value: 4},
		{},
		{5, nil, nil},
	}

	fmt.Println(root, root1, nodes)

	fmt.Println("\n----------")

	root1.print()

	fmt.Println("\n----------")

	root1.right.left.print()
	fmt.Print(" ")
	root1.right.left.setValue(4)
	root1.right.left.print()
	fmt.Print(" ")
	root1.right.left.setValue2(4)
	root1.right.left.print()

	fmt.Println("\n----------")

	root_copy := &root1
	root_copy.right.left.print()
	fmt.Print(" ")
	root_copy.right.left.setValue2(6)
	root1.right.left.print()

}
