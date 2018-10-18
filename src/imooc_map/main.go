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


	fmt.Println("======================\ntry to use golang to find the longest non-repetitive string in the map\n=================")

	var data = []string {"fdsafdsafdsaf", "fdsafkjdslaf", "fdjsaofeoih", "fdsafgbffa", "vcxzewfqfad", "fdjksafnjhowefn", "fdnkslavndsle"}
	result := make(map[int][]string)

	for _, v := range data {
		str := []rune(v)

		temp := make(map[rune]uint8)

		for _, c := range str {
			if _, ok := temp[c]; ok {
				temp[c]++
			} else {
				temp[c] = 1
			}
		}

		result[len(temp)] = append(result[len(temp)], v)
	}

	fmt.Println(result)

	fmt.Println("======================\ncopy from imooc\n=================")
	var data2 = []string {"112341516"}
	// fdjksafnjhowefn  => fdjksanhowe

	for _, v := range data2 {
		fmt.Println(v, lengthOfNonRepeatingSubstr(v))
	}
}

func lengthOfNonRepeatingSubstr(s string) int {
	fmt.Println("↓↓↓↓")

	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		fmt.Println("lastOccurred[ch]", lastI, ok)
		fmt.Println("start: ", start, "\tmaxLength: ", maxLength, "\ti: ", i, "\tch: ", string(ch))

		if ok && lastI >= start {
			start = lastI + 1
		}

		if lastOccurred[ch] >= start {
			start = lastOccurred[ch] + 1
		}

		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i

		fmt.Println("start: ", start, "\tmaxLength: ", maxLength, "\n")
	}

	fmt.Println("↑↑↑↑")
	return maxLength
}
