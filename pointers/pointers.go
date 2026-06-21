package main

import "fmt"

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println(num)
// }

func changeNum(num *int) {
	*num = 5
	fmt.Println(*num)
}
func main() {
	var num int = 1
	changeNum(&num)
	fmt.Println("after change", num)
	fmt.Println(&num)
}
