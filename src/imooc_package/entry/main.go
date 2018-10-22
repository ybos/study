package main

// 目录的名字
import "imooc_package"

func main() {
	// 包的名字
	node := imooc_package_node.CreateNode(1)

	node.Left = imooc_package_node.CreateNode(2)
	node.Right = imooc_package_node.CreateNode(3)

	node.Print()
	node.Left.Print()
	node.Right.Print()
}
