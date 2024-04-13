package main


import (
	"fmt"
	"sync"
	"time"
)

type Councurrency struct {
	mu   sync.Mutex
	data []int
	sum  int
	wg sync.WaitGroup
}

func NewCouncurrency(array []int) *Councurrency {
	c := Councurrency{
		data: make([]int, len(array)),
	}
	for i, value := range array {
		c.data[i] = value
	}
	return &c
}

func (c *Councurrency) Step(index int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Printf("Start %d\n", index)
	c.sum += c.data[index] * c.data[index]
	time.Sleep(1 * time.Second)
	fmt.Printf("End %d\n", index)
	c.wg.Done()
}

func main() {
	sl := []int{2, 4, 6, 8, 10}
	c := NewCouncurrency(sl)
	
	for i := range sl {
		c.wg.Add(1)
		go c.Step(i)
	}
	c.wg.Wait()
	fmt.Println(c.sum)
}