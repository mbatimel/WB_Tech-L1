package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
// 1 способ: Завершение канала, который используется горутиной, 
//но чтобы не схватить deadlock, мы синхронизируем с помощью sync.WaitGroup
ch := make(chan int)
wg := new(sync.WaitGroup)
wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			value, ok := <-ch
			if !ok {
				fmt.Println("[Goroutine 1]: Done")
				return
			}
			fmt.Println(value)
		}
	}()
time.Sleep(1*time.Second)
ch <- 1
time.Sleep(1*time.Second)
ch <- 2	
time.Sleep(1*time.Second)
ch <- 3
time.Sleep(1*time.Second)
close(ch)
wg.Wait()

// 2 способ: Мы сигнализируем о том что программа завершается. Для этого мы используем for-select.
// создается специальный канал, который будет жать сигнала остановки, то есть ctl+c

//канал который будет ожидать закрытия
stopChan := make(chan struct{})
//канал который закроется после закрытия основного канала
stoppedChan := make(chan struct{})
	go func(){
		defer close(stoppedChan)

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case now := <-ticker.C:
				fmt.Printf("Goroutine: tick %s\n", now.UTC().Format("20060102-150405.000000000"))
			case <-stopChan:
				fmt.Println("[Goroutine 2]: Done")
				return
			}
		}
	}()
c := make(chan os.Signal, 1)
signal.Notify(c, os.Interrupt)
<-c
fmt.Println("\nmain: received C-c - shutting down")
fmt.Println("main: telling goroutines to stop")
close(stopChan)
<-stoppedChan


// 3 способ: Использование контекстного пакета
ctx, cancel := context.WithCancel(context.Background())
ch3 := make(chan int)	
wg.Add(1)
go func() {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[Goroutine 3]: Done")
			return
		default:
			value := <-ch3
			fmt.Println(value)
		}
	}
}()
time.Sleep(time.Second)
ch3 <- 1
time.Sleep(time.Second)
ch3 <- 2
time.Sleep(time.Second)
ch3 <- 3
cancel()
close(ch3)
wg.Wait()
}