package main

import "fmt"

func findFirst(nums []int, target int) int {
	idx := -1
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] >= target {
			end = mid - 1
		} else {
			start = mid + 1
		}
		if nums[mid] == target {
			idx = mid
		}
	}
	return idx
}

func findLast(nums []int, target int) int {
	idx := -1
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
		if nums[mid] == target {
			idx = mid
		}
	}
	return idx
}

func main() {
	arr := []int{-1, -1, 2, 2, 2, 2, 2, 2, 9}
	t := 2
	a := findFirst(arr, t)
	b := findLast(arr, t)
	fmt.Println(a, b)
}
