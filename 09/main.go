package main

import (
	"fmt"
	"sync"
)
/*
если мы говорим о конвеере, то получается мы должны последовательно выполнять задачи:
1) запись числа в канал 1
2) возведение в квадрат числа и запись его во 2 канал
3) вывод числа
воспользуемся кокурентным воздействием над каждым из трех дейсткий,
синхранизируем действие каждой функции с помощь sync.WaitGroup
*/
func main() {
	var wg sync.WaitGroup

	ch1 := make(chan int)
	ch2 := make(chan int)
	data := []int{1, 2, 3, 4, 5, 6}

	wg.Add(1)
	go func() {
		
		defer wg.Done()
		for num := range ch1 {
			//fmt.Printf("add to ch2: %d^2\n", num)
			ch2 <- num * num
		}
		close(ch2)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch2 {
			fmt.Printf("result: %d\n", num)
		}
		fmt.Println()
	}()
	
	for _, value := range data {
		//fmt.Printf("add to ch1: %d\n", value)
		ch1 <- value
	}


	
	close(ch1)
	wg.Wait()
}