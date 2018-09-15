package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s))
	fmt.Println(cap(s))
	ss := s[1:3]        //容量只看前面的数字
	fmt.Println(len(ss))
	fmt.Println(cap(ss))

	ss = append(ss, 4)
	for _, v := range ss {
		v = v + 10
		fmt.Println(v)
	}
	fmt.Println(ss)

	for i := range ss {
		fmt.Println(i)
		ss[i] = ss[i] + 10
	}
	fmt.Println(ss)
	fmt.Println(s)
}
