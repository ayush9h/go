package main

import "fmt"

// iterating
func main() {

	nums := []int{1, 2, 3, 5, 6}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		fmt.Println(nums[i])
	}
	fmt.Println(sum)

	// using range(idx, element)
	for i, num := range nums {
		fmt.Println(i, num)
		// fmt.Println(num)
	}

}
