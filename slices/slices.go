package main

import (
	"fmt"
	"slices"
)

func main() {
	// data-type, initial size, capacity
	var nums = make([]int, 2, 3)

	nums = append(nums, 3)
	nums = append(nums, 3)
	nums = append(nums, 3)
	nums = append(nums, 3)
	nums = append(nums, 3)

	nums = append(nums, 3)
	nums = append(nums, 3)
	nums = append(nums, 3)
	nums = append(nums, 3)
	fmt.Println(nums)
	fmt.Println(cap(nums))

	// copy(dst,src)
	//
	// slice operator
	var nums1 = []int{1, 2, 3}
	fmt.Println(nums1[:1])

	//slices
	var num1 = []int{1, 2, 3}
	var num2 = []int{1, 2, 3}
	fmt.Println(slices.Equal(num1, num2))
}
