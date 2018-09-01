package main

import "fmt"

func main() {
	i := 1
	for i <= 6666 {
		fmt.Println(i)
		i += 1
	}

	for j := 2; j < 8888; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for m := 0; m <= 100; m++ {
		if m%3 == 0 {
			continue
		}
		fmt.Println(m)
	}
}
