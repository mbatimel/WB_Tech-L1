package main

import (
	"fmt"
	"time"
)

func deltatime(startTime *time.Time) int {
	return int(time.Since(*startTime).Seconds())
}
/*
	последовательно отправляем в горутину значение со счетчика
	 и сравниваем значение  которое мы определил со временем начала запуска программы
*/
func main() {
	n := 0
	fmt.Printf("Enter time: ")
	fmt.Scan(&n)
	startTime := time.Now()
	ch := make(chan int)

	go func() {
		for data := range ch {
			fmt.Println(data)
		}
	}()

	for i := 0; deltatime(&startTime) <= n; i++ {
		ch <- i
	}
	close(ch)
}