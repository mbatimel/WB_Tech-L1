package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	list list.List
	len  int
}

func (q *Stack) Init() {
	q.list.Init()
}

func (q *Stack) Push(el interface{}) {
	q.list.PushFront(el)
	q.len++
}

func (q *Stack) Top() (interface{}, bool) {
	if q.len == 0 {
		return 0, false
	}
	return q.list.Front().Value, true
}

func (q *Stack) Pop() {
	if q.len == 0 {
		return
	}
	q.list.Remove(q.list.Front())
	q.len--
}

func main() {
	var q Stack
	q.Init()
	q.Push("asd")
	q.Push(1)
	fmt.Println(q.Top())
	q.Pop()
	fmt.Println(q.Top())
	q.Pop()
	fmt.Println(q.Top())
}

// 1 true
// asd true
// 0 false