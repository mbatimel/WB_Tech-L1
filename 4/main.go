package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)
type Event struct {
	 wg sync.WaitGroup
	 mu sync.Mutex
	 ch chan int64
}
func main() {
	ev := new(Event)
	workers := 0 
	fmt.Printf("Enter the number of workers: ")
	fmt.Scan(&workers)
	
	ev.ch = make(chan int64, workers)
	for i := 0; i < workers; i++ {
		ev.wg.Add(1)
		
		go func(i int){
			defer ev.wg.Done()
			for n := range ev.ch {
					ev.mu.Lock()
					fmt.Printf("Worker %d : %d\n",i,n)
					ev.mu.Unlock()
					time.Sleep(10 * time.Second)
				}
				fmt.Printf("Worker %d : Done\n", i)
		}(i)
		
	}
	go ev.check_end()
	for {
		n := rand.Int63()
		data := n
		/*	вторая реализация, для ручного ввода данных
			data := ""
			fmt.Printf("Enter data: ")
			fmt.Scan(&data)
		*/
		fmt.Printf("Entered data:	%d\n", data)
		ev.ch <- data
	}

}
func (ev *Event) check_end() {
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("\r")
		close(ev.ch)
		ev.wg.Wait()
		os.Exit(0)
	}()
}