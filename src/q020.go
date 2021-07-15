package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	data    []int
	minData []int
	lock    sync.RWMutex
}

func (s *Stack) Push(item int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = append(s.data, item)
	if len(s.minData) == 0 {
		s.minData = append(s.minData, item)
	} else {
		if s.minData[len(s.minData)-1] >= item {

			s.minData = append(s.minData, item)
		}
	}
}
func (s *Stack) Pop() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.minData) == 0 {
		return 0
	} else {
		if s.minData[len(s.minData)-1] == s.data[len(s.data)-1] {
			s.minData = s.minData[:len(s.minData)-1]
		}
		tmp := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return tmp
	}
}
func (s *Stack) Min() int {
	return s.minData[len(s.minData)-1]
}

func main() {
	s := &Stack{}
	s.Push(10)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(2)
	fmt.Println(s.Pop())
	fmt.Println("111123123", s.Min())
	fmt.Println(s.Pop())
	fmt.Println("111123123", s.Min())
	fmt.Println(s.Pop())
	fmt.Println("111123123", s.Min())

}
