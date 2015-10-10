package deadpool

import (
	"errors"
)

type DeadPool struct {
	Threads int       //number of threads you want the pool to have
	sem     chan bool //a channel that limits the number
}

func NewPool(threads int) (*Deadpool, error) {
	if threads < 1 {
		return errors.New("You cannot pass a number less 1")
	}
	sem := make(chan bool, threads)
	return &DeadPool{Threads: threads, sem: sem}
}

func (d *DeadPool) Add(fn func(args ...interface{})) {
}

func (d *DeadPool) Wait() {
}
