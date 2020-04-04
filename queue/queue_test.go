package queue

import (
    "testing"
    "../util"
)

func TestEnqueue(t *testing.T) {
    var b4, after interface{}
    for iterations := 0; iterations < 50; iterations++ {
        newQ := &Queue{}
        for loops := 0; loops < 10; loops++ {
            random := util.GetRandIntn(loops*3)

            if newQ.Size() > 1 {
                b4 = newQ.Tail
            } else if newQ.Head != nil {
                b4 = newQ.Head
            }

            if newQ.Enqueue(random) {
                if newQ.Size() > 1 {
                    after = newQ.Tail
                } else {
                    after = newQ.Head
                }

                if b4 == after {
                    t.Errorf("Did not append %v to list last %+v, newlyAdded %+v", random, b4, after)
                }
            }
        }
    }
}

func TestDequeue(t *testing.T) {
    var newHead interface{}
    for iterations := 0; iterations < 50; iterations++ {
        newQ := &Queue{}
        newQ.FillList(100)
        for loops := 0; loops < 10; loops++ {
            oldHead := newQ.Head
            newHead = newQ.Dequeue()

            if oldHead.Data != newHead {
                t.Errorf("Did not remove right data expected %+v, but got %+v", oldHead, newHead)
            }
        }
    }
}

func TestPeek(t *testing.T) {
    for iterations := 0; iterations < 50; iterations++ {
        newQ := &Queue{}
        newQ.FillList(100)
        for loops := 0; loops < 10; loops++ {
            oldHead := newQ.Peek()
            oldVal := newQ.Dequeue()
            currFront := newQ.Peek()
            
            if oldHead != oldVal || currFront != newQ.Head.Data {
                t.Errorf("Not peeking right expected: %+v, got: %v", newQ.Head, currFront)
            }
        }
    }
}
