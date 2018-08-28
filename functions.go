package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	result1 := plus(66, 88)
	fmt.Println(result1)

	result2 := plusPlus(100, 200, 300)
	fmt.Println(result2)
}
