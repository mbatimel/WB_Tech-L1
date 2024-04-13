package main

import (
	"fmt"
	"sync"
)



func conversion_channel(nums []int) <-chan int {
    out := make(chan int, len(nums))
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
    square_chan2 := square_nums(c)
    for n := range merge(square_chan, square_chan2) {
        fmt.Println(n) 
        }



}

func square_nums(c <-chan int) <-chan int {
    mu := new(sync.Mutex)
    out := make(chan int, len(c))
    go func(){
       mu.Lock()
        for n := range c {
            out <- n*n
        }
        mu.Unlock()
        close(out)
    }()
    return out
}
func merge(in ...<-chan int) <-chan int{
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