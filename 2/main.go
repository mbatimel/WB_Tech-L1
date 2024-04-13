package main

import (
	"fmt"
	"sync"
	"time"
)



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