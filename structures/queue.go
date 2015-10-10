package structures

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

func (q *FuncQueue) Enqueue(fn func()) {
	NewNode := NewFuncNode(fn)
	if q.front == nil {
		q.front = New
		q.last = New
	} else {
		q.last.Next = New
		q.last = NewNode
	}
}

// the thread pool just needs a function, so thats what we're gonna give it
func (q *Dequeue) Dequeue() func() {
	prev = nil
	temp = q.front
	target = q.last.F
	for {
		if temp == nil {
			break
		}
		if temp.F == target {
			q.last = prev
			break
		}
		prev = temp
		temp = temp.Next
	}
	return target
}
