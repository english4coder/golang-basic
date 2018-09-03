package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key1"] = 10
	m["key2"] = 20
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	v1 := m["key1"]
	fmt.Println("v1:", v1)

	delete(m, "key2")
	fmt.Println("map:", m)

	value, present := m["key1"]
	fmt.Println("value/present", value, present)

	_, present2 := m["key2"]
	fmt.Println("present2", present2)

	n := map[string]int{"addr1": 1, "addr2": 2}
	fmt.Println("map:", n)
}
