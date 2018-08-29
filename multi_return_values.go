package main

import "fmt"

func vals() (int, int) {
	return 8, 6
}

func main() {
	a, b := vals()
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	_, c := vals()
	fmt.Println("c:", c)

	d, _ := vals()
	fmt.Println("d:", d)
}
