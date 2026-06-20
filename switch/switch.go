package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5
	switch i {
	case 1:
		fmt.Println("1")

	case 2:
		fmt.Println("2")
	default:
		fmt.Println("3")
	}

	//multiple condition
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its weekend")
	default:
		fmt.Println("its a bloody workday")
	}

	//type switch - making function switch
	whoamI := func(i any) {
		switch t := i.(type) {
		case int:
			fmt.Println("this is a integer")
		case bool:
			fmt.Println("this is a boolean")
		default:
			fmt.Println("this is something fishy", t)
		}
	}
	whoamI("myname")
}
