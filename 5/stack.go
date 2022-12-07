package main

import (
	"sync"
)

type Stack struct {
	dll   []string
	mutex sync.Mutex
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(x string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.dll = append(s.dll, x)
}

func (s *Stack) Pop() (string, bool) { // ok idiom
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.dll) == 0 {
		return "", false
	}
	res := (s.dll)[len(s.dll)-1]
	s.dll = (s.dll)[:len(s.dll)-1]
	return res, true
}

func (s *Stack) Reverse() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, j := 0, len(s.dll)-1; i < j; i, j = i+1, j-1 {
		s.dll[i], s.dll[j] = s.dll[j], s.dll[i]
	}
}
