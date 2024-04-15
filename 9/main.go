package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch1 := make(chan int)
	ch2 := make(chan int)
	data := []int{1, 2, 3, 4, 5, 6}


	go func() {
		wg.Add(1)
		defer wg.Done()
		for num := range ch1 {
			fmt.Printf("add to ch2: %d^2\n", num)
			ch2 <- num * num
		}
		close(ch2)
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		for num := range ch2 {
			fmt.Printf("result: %d\n", num)
		}
		fmt.Println()
	}()
	
	for _, value := range data {
		fmt.Printf("add to ch1: %d\n", value)
		ch1 <- value
	}


	
	close(ch1)
	wg.Wait()
}