package stack

import (
    "../queue"
)

type MyStack interface {
    Push(el interface{}) bool
    Pop() interface{}
    Peek() interface{}
    Size() int
    Capacity() int
}

type Stack struct {
    items *queue.DeqLL
}

func (s *Stack) Push(el interface{}) bool {
    return s.items.AddFront(el)
}

func (s *Stack) Pop() interface{} {
    return s.items.RemoveFront()
}

func (s *Stack) Peek() interface{} {
    return s.items.PeekFront()
}

func (s *Stack) Size() int {
    return s.items.Size()
}

func NewStack() *Stack {
    return &Stack{ queue.NewDeqLL() }
}
