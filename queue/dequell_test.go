package queue

import (
    "testing"
)

func TestAddFrontLL(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeq := NewDeqLL(iterations)
        currVal := iterations*2
        oldSize := newDeq.Size()
        if newDeq.AddFront(currVal) {
            curr := newDeq.PeekFront()
            newSize := newDeq.Size()
            
            if curr != currVal {
               t.Errorf("Expected %v, but got %v, or incorrect size old %d, new %d", currVal, curr, oldSize, newSize)
            }
        }
    }
}

func TestAddBackLL(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeq := NewDeqLL(iterations)
        currVal := iterations*2
        oldSize := newDeq.Size()
        if newDeq.AddBack(currVal) {
            curr := newDeq.PeekBack()
            newSize := newDeq.Size()
            
            if curr != currVal {
               t.Errorf("Expected %v, but got %v, or incorrect size old %d, new %d", currVal, curr, oldSize, newSize)
            }
        }
    }
}

func filledDeqLL(delta int) *DeqLL {
    loops := delta+50
    newDeq := NewDeqLL(loops)
    for loops != 0 {
        newDeq.AddBack(loops)
        loops--
    }

    return newDeq
}

func TestRemoveFrontLL(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeq := filledDeqLL(iterations*5)
        for loops := 0; loops < iterations*5; loops++ {
            oldTop := newDeq.PeekFront()
            oldSize := newDeq.Size()
            removed := newDeq.RemoveFront()
            newVal := newDeq.PeekFront()
            newSize := newDeq.Size()

            if removed != oldTop || newSize == oldSize {
                t.Errorf("Did not remove %v oldTop%p %+v, newVal%p %+v or incorrect size old %d, new %d", removed, oldTop, oldTop, newVal, newVal, oldSize, newSize)
            }
        }
    }
}



func TestRemoveBackLL(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeq := filledDeqLL(iterations*5)
        for loops := 0; loops < iterations*5; loops++ {
            oldTop := newDeq.PeekBack()
            oldSize := newDeq.Size()
            removed := newDeq.RemoveBack()
            newVal := newDeq.PeekBack()
            newSize := newDeq.Size()

            if removed != oldTop || oldTop == newVal {
                t.Errorf("Did not remove %v oldTop%p %+v, newVal%p %+v or incorrect size old %d, new %d", removed, oldTop, oldTop, newVal, newVal, oldSize, newSize)
            }
        }
    }
}

func TestPeekFrontLL(t *testing.T) {
    newDeq := NewDeqLL(100)
    for iterations := 1; iterations < 100; iterations++ {
        newDeq.AddFront(iterations)
        currTop := newDeq.PeekFront()
        if currTop !=  iterations {
            t.Errorf("Expected %+v, but got %v", iterations, currTop)
        }
    }
}

func TestPeekBackLL(t *testing.T) {
    newDeq := NewDeqLL(100)
    for iterations := 1; iterations < 100; iterations++ {
        newDeq.AddBack(iterations)
        currTail := newDeq.PeekBack()
        if currTail !=  iterations {
            t.Errorf("Expected %+v, but got %v", iterations, currTail)
        }
    }
}
