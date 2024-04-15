package main

import (
	"fmt"
	"sync"
)

type SomeMap struct {
	sync.RWMutex
	m map[int]int

}

func main(){
	wg := sync.WaitGroup{}
	someMap := &SomeMap{
		m: map[int]int{},
	}
	for i:=0;i<100;i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			someMap.Lock()
			defer someMap.Unlock()
			someMap.m[i] = i
		}(i)
	}
	for i:=0;i<100;i++{
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			someMap.RLock()
			defer someMap.RUnlock()
			if number, ok := someMap.m[i]; ok {
				fmt.Printf("Map key: %d value: %d\n",i,number)
			}else{
			fmt.Printf("Number does not exists\n")
		}
		}(i)
	}
	wg.Wait()
}