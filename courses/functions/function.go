package main

import "fmt"

//	func add(a int, b int) int {
//		return a + b
//	}
func add(a, b int, c, d bool) (int, bool) {
	return a + b, c && d
}

func variadic(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}

func main() {
	result1, result2 := add(2, 3, true, true)
	fmt.Println(result1, result2)

	result := variadic(3, 5, 4, 5, 6, 1, 3, 3)
	fmt.Println(result)
}
