package main

import (
	"fmt"
	"sync"
)

func checkNums(i int, wg *sync.WaitGroup) {
	defer wg.Done() // runs after completion of the function with defer keyword
	fmt.Println(i)
}

func main() {

	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go checkNums(i, &wg)
	}

	wg.Wait()
	// Bad Habit
	// time.Sleep(time.Second * 2)

}
