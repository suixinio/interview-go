package main

import (
	"fmt"
	"time"
)

type Semaphore struct {
	Channel chan int
}

func NewSemaphore(parallelism int) *Semaphore {
	return &Semaphore{Channel: make(chan int, parallelism)}
}

func (s Semaphore) Lock() {
	s.Channel <- 1
}
func (s Semaphore) Unlock() {
	<-s.Channel
}

func main() {
	//	共享锁
	sem := NewSemaphore(5)
	i := 0
	for {
		sem.Lock()
		i++
		go func() {
			defer sem.Unlock()
			time.Sleep(2 * time.Second)
			fmt.Println(i)
			i--
		}()
	}
}
