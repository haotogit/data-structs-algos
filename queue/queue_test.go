package queue

import (
    "testing"
    "../util"
)

func TestEnqueue(t *testing.T) {
    for iterations := 0; iterations < 50; iterations++ {
        newQ := NewQ()
        for loops := 0; loops < 10; loops++ {
            random := util.GetRandIntn(loops*3)
            oldSize := newQ.Size()
            if newQ.Enqueue(random) {
                lastItem := newQ.currDeq.PeekBack()
                newSize := newQ.Size()

                if oldSize == newSize || random != lastItem {
                    t.Errorf("Did not append %v to list, last item is %v, or size is wrong oldSize %d newSize %d", random, lastItem, oldSize, newSize)
                }
            }
        }
    }
}

func TestDequeue(t *testing.T) {
    for iterations := 0; iterations < 50; iterations++ {
        newQ := NewQ()
        newQ.currDeq.FillList(10*iterations)
        oldSize := newQ.Size()
        for loops := 0; loops < 10; loops++ {
            oldHead := newQ.Peek()
            newHead := newQ.Dequeue()
            newSize := newQ.Size()

            if oldHead != newHead || oldSize == newSize {
                t.Errorf("Did not remove right data expected %+v, but got %+v, or size wrong old %d, new %d", oldHead, newHead, oldSize, newSize)
            }
        }
    }
}

func TestPeek(t *testing.T) {
    for iterations := 0; iterations < 50; iterations++ {
        newQ := NewQ()
        newQ.currDeq.FillList(10*iterations)
        for loops := 0; loops < 10; loops++ {
            currFront := newQ.Peek()
            deqFront := newQ.currDeq.PeekFront()
            
            if currFront != deqFront {
                t.Errorf("Not peeking right expected: %+v, got: %v", deqFront, currFront)
            }
        }
    }
}
