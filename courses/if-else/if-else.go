package main

import "fmt"

func main() {

	var admin bool = true
	var adminName string = "Foo"

	// Conditional operators if-else
	if admin && len(adminName) >= 3 {
		fmt.Println("Logged in successfully")
	}

	// declare variables in if-else block
	if age := 18; age >= 15 {
		fmt.Println("he is adult")
	} else if age >= 12 {
		fmt.Println("he is teenager")
	}
}
