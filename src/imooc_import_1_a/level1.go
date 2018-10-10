package imooc_import_1_a

import (
	"fmt"
	"imooc_import_2_a"
)

func init() {
	fmt.Println("init Level-1-a")
}

func ShowLevel1a() {
	fmt.Println("show Level-1-a")
	imooc_import_2_a.ShowLevel2a()
}