package main

import (
	. "fmt"
)

func main() {
	// 声明并初始化一个映射
	// 使用make来创建映射
	map1 := make(map[string]string)

	map1["John"] = "boy"
	map1["Neil"] = "cute boy"
	map1["Sofia"] = "girl"

	for k, v := range map1 {
		Println(k, " is a ", v)
	}

	Println("\r\n")

	// 使用map来创建映射
	map2 := map[string]string{"red": "#da1337", "orange": "#ff6600", "black": "#000"}

	for k, v := range map2 {
		Println(k, " is a ", v)
	}

	Println("\r\n")

	// 初始化一个空的映射
	map3 := map[string]int{}

	map3["John"] = 15
	map3["Neil"] = 27

	for k, v := range map3 {
		Println(k, " is a ", v)
	}

	// 映射的读取
	// 读取某个键的映射值时，返回有两个结果，第一个是值，第二个是键是否存在的布尔值
	map4 := map[string]int{"John": 10, "Neil": 27, "Todd": 0}
	value, exists := map4["Neil"]

	if exists {
		Println("map4 has key: Neil, value is: ", value)
	} else {
		Println("map4 doesn't have key: Neil")
	}

	value, exists = map4["Sofia"]

	if exists {
		Println("map4 has key: Sofia, value is: ", value)
	} else {
		Println("map4 doesn't have key: Sofia")
	}

	// 也可以选择抛弃不使用第二个返回值
	value = map4["Sofia"]

	if value == 0 {
		Println("map4 doesn't have key: Sofia")
	} else {
		Println("map4 has key: Sofia, value is: ", value)
	}

	// 但是如果值就是0，这时候的判断就会存在歧义
	value = map4["Todd"]

	if value == 0 {
		Println("map4 doesn't have key: Todd, but map4 has key 'Todd' and Todd is 0 now")
	} else {
		Println("map4 has key: Todd, value is: ", value)
	}

	Println("\r\n")

	// 删除某个键值对
	map5 := make(map[string]int)

	map5["John"] = 10
	map5["Neil"] = 27
	map5["Sofia"] = 15

	for k, v := range map5 {
		Println("key: ", k, ", value: ", v)
	}

	delete(map5, "John")

	for k, v := range map5 {
		Println("key: ", k, ", value: ", v)
	}

	Println("\r\n")

	// 查看映射在参数传递时的使用方式
	map6 := map[string]int{"John": 15, "Neil": 27, "Sofia": 20}

	Printf("map6 points to: %p\n", &map6)
	for k, v := range map6 {
		Println("key: ", k, ", value: ", v)
	}

	deleteOne(map6, "Sofia")

	for k, v := range map6 {
		Println("key: ", k, ", value: ", v)
	}
}

func deleteOne(item map[string]int, key string) {
	Printf("item points to: %p\n", &item)

	delete(item, key)
}
