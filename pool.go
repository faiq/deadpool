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
	d.sem <- true //try to put something in the buffered channel, if its full it's going to block
	go func() {
		defer func() { <-sem }() //read from the channel to free it up
		fn(args)
	}()
}

func (d *DeadPool) Wait() {
	for i := 0; i < cap(sem); i++ {
		sem <- true //pushes trues back on once all the go routines start finishing up
	}
}
