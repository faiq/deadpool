package main

import (
	"fmt"
)

type FuncNode struct {
	F    func() //pointer to the function to exec
	Next *FuncNode
}

func NewFuncNode(f func()) *FuncNode {
	return &FuncNode{f, nil}
}

type FuncQ struct {
	front *FuncNode
	last  *FuncNode
}

func NewFuncQueue() *FuncQ {
	return &FuncQ{}
}

func (q *FuncQ) Enqueue(fn func()) {
	NewNode := NewFuncNode(fn)
	if q.front == nil {
		q.front = NewNode
		q.last = NewNode
	} else {
		q.last.Next = NewNode
		q.last = NewNode
	}
}

// the thread pool just needs a function, so thats what we're gonna give it
func (q *FuncQ) Dequeue() func() {
	var prev *FuncNode
	temp := q.front
	target := &q.last.F
	ret := q.last.F
	for {
		if temp == nil {
			break
		}
		if &(temp.F) == target {
			q.last = prev
			break
		}
		prev = temp
		temp = temp.Next
	}
	return ret
}

func main() {
	x := NewFuncQueue()
	x.Enqueue(func() {
		fmt.Printf("ayy lmao")
	})
	y := x.Dequeue()
	y()
}
