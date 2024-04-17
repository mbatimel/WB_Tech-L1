package main

import (
	"fmt"
	"sync"
	"time"
)
	/*
        Изначально предполагалось что я использую парттерн Fan-out, fan-in для конкурентного
        возведения в степень, но я немного не понял работу, поэтому убрал один канал из реализации
        и использовал его вместо двух  
		Создаем цикл по массиву
		В данном цикле происходит 5 итераций, соответственно кол-ву чисел в массиве
		Далее создается 5 горутин, в которых присутствует метод Done, сообщающий,
		что выполнилась задача в WaitGroup.
		После того, как все горутины были созданы, мы выходим из цикла
		и попадаем на метод Wait, который в свою очередь ждет пока выполнятся все задачи из wg
		То есть ждет пока все числа из массива с помощью math.Pow возведутся в степень
		и выведутся в stdout. Затем, wg закрывается"
	*/

// функция которая переводи массив в канал
func conversion_channel[T any](nums []T) <-chan T {
    out := make(chan T, len(nums))
    go func(){
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func main() {
    a:= []int{2,4,6,8,10}
    c:= conversion_channel(a)
    square_chan := square_nums(c)

    for n := range merge(square_chan) {
        fmt.Println(n) 
        }



}
// возводим в квадрат все числа из канала 
func square_nums(c <-chan int) <-chan int {
    out := make(chan int, len(c))
   
    go func(){
        for n := range c {
            fmt.Println("Start num: ", n)
            out <- n*n
            time.Sleep(1 * time.Second)
            
        }
        
        close(out)
    }()
    return out
}
// переводим  все в канал для вывода полученных значений
func merge(in ... <-chan int) <-chan int{
    var wg sync.WaitGroup
    out := make(chan int)

    output := func(c <-chan int) {
        for n := range c {
        out <- n
        }
        wg.Done()
    }
    wg.Add(len(in))
    for _, c := range in {
        go output(c)
    }

    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
