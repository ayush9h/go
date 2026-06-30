package main

import "fmt"

func main() {
	var nums [4]int
	nums[0] = 1
	fmt.Println(len(nums))

	fmt.Println(nums)

	nums2 := [2][2]int{{2, 2}, {1, 1}}
	fmt.Println(nums2)
}
