package main

import "fmt"

func main() {
	m1 := map[string]string {
		"name": "Joker",
		"course": "iMooc",
		"age": "28",
		"sex": "male",
	}

	m2 := make(map[string]string) // m2 == empty map
	var m3 map[string]string // m3 == nil

	fmt.Println(m1, m2, m3)

	for k, v := range m1 {
		fmt.Println("loop m1: ", k, v)
	}

	fmt.Println("Getting values")
	fmt.Printf("get name: %s\n", m1["name"])
	fmt.Printf("get a inexistent mane: %s\n", m1["mane"])

	if name, ok := m1["mane"]; ok {
		fmt.Printf("Key mane exists: %s\n", name)
	} else {
		fmt.Println("Key mane inexistent.")
	}

	fmt.Println("Deleting values")
	name, ok := m1["name"]
	fmt.Println(name, ok, m1)

	delete(m1, "name")
	fmt.Println(name, ok, m1)
}
