package main

import "fmt"

func main() {
	str := "sadlfjlsdjflsdjf"
	r := []rune(str)
	copy(r[4:4+3], []rune("abc"))
	fmt.Printf("Before: %s\n", str)
	fmt.Printf("After: %s\n", string(r))
}
