package main

import (
	"fmt"
	"math"
)

const str = "constantdemo"

func main() {
	fmt.Println(str)
	const num1 = 8888888
	const num2 = 666666666 / num1
	fmt.Println(num2)
	fmt.Println(int64(num2))
	fmt.Println(math.Sin(num2))
}
