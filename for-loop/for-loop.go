package main

import "fmt"

// only construct in go loop - for
func main() {
	i := 1
	//while loop
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//classic loop
	for i := 1; i <= 10; i++ {

		//break the execution of for loop
		if i == 8 {
			break
		}
		//breaks the current execution of for loop
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
}
