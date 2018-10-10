package imooc_import_1_b

import (
	"fmt"
	"imooc_import_2_b"
)

func init() {
	fmt.Println("init Level-1-b")
}

func ShowLevel1b() {
	fmt.Println("show Level-1-b")
	imooc_import_2_b.ShowLevel2b()
}