package stack

import (
    "../queue"
)

type MyStack interface {
    Push(el interface{}) bool
    Pop() interface{}
    Peek() interface{}
    Size() int
}

type Stack struct {
    currStack *queue.Deq
}

func (s *Stack) Push(el interface{}) bool {
    return s.currStack.AddFront(el)
}

func (s *Stack) Pop() interface{} {
    return s.currStack.RemoveFront()
}

func (s *Stack) Peek() interface{} {
    return s.currStack.PeekFront()
}

func (s *Stack) Size() int {
    return s.currStack.Size()
}

func newStack() *Stack {
    return &Stack{ queue.NewDeq() }
}
