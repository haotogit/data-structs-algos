package stack

import (
    "testing"
)

func TestPush(t *testing.T) {
    for iterations := 0; iterations < 100; iterations++ {
        newStack := NewStack()
        for loops := 0; loops < 10; loops++ {
            newVal := loops+iterations
            oldSize := newStack.Size()

            if newStack.Push(newVal) {
                newSize := newStack.Size()
                if newStack.Peek() != newVal || oldSize == newSize {
                    t.Errorf("Did not push to stack or incorrect size old %d, new %d", oldSize, newSize)
                }
            }
        }
    }
}

func TestPop(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newStack := NewStack()
        newStack.FillList(iterations*5)
        for loops := 0; loops < iterations*5; loops++ {
            oldSize := newStack.Size()
            oldTop := newStack.Peek()
            popped := newStack.Pop()
            newSize := newStack.Size()
            if oldTop != popped || oldSize == newSize {
                t.Errorf("Did not pop right, expected %v, but got %v, or incorrect size old %d, new %d", oldTop, popped, oldSize, newSize)
            }
        }
    }
}

func TestPeek(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newStack := NewStack()
        newStack.Push(iterations)
        peeked := newStack.Peek()
        if peeked != iterations {
            t.Errorf("Incorrect peeking, expcted %v but got %v", iterations, peeked)
        }
    }
}
