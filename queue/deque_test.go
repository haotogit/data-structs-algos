package queue

import (
    "testing"
)

func TestAddFront(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeq := NewDeq(iterations)
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

func TestAddBack(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeq := NewDeq(iterations)
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

func TestIncCap(t *testing.T) {
    for iterations := 1; iterations < 200; iterations++ {
        newDeq := filledDeq(iterations)
        if iterations % 2 == 0 {
            newDeq.AddFront(iterations)
        } else {
            newDeq.AddBack(iterations)
        }

        if newDeq.Capacity() != (iterations+1)*2 {
            t.Errorf("Did not increase capacity expected %d, but got %d", (iterations+1)*2, newDeq.Capacity())
        }
    }
}

func filledDeq(delta int) *Deq {
    loops := delta
    newDeq := NewDeq(loops)
    for loops != 0 {
        newDeq.AddBack(loops)
        loops--
    }

    return newDeq
}

func TestRemoveFront(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeq := filledDeq(iterations*5)
        for loops := 0; loops < iterations*5; loops++ {
            oldTop := newDeq.PeekFront()
            oldSize := newDeq.Size()
            removed := newDeq.RemoveFront()
            newVal := newDeq.PeekFront()
            newSize := newDeq.Size()

            if removed != oldTop || newSize == oldSize {
                t.Errorf("Did not remove %v oldTop%p %+v, currFront %p %+v or incorrect size old %d, new %d", removed, oldTop, oldTop, newVal, newVal, oldSize, newSize)
            }
        }
    }
}



func TestRemoveBack(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeq := filledDeq(iterations*5)
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

func TestPeekFront(t *testing.T) {
    newDeq := NewDeq(100)
    for iterations := 1; iterations < 100; iterations++ {
        newDeq.AddFront(iterations)
        currTop := newDeq.PeekFront()
        if currTop !=  iterations {
            t.Errorf("Expected %+v, but got %v", iterations, currTop)
        }
    }
}

func TestPeekBack(t *testing.T) {
    newDeq := NewDeq(100)
    for iterations := 1; iterations < 100; iterations++ {
        newDeq.AddBack(iterations)
        currTail := newDeq.PeekBack()
        if currTail !=  iterations {
            t.Errorf("Expected %+v, but got %v", iterations, currTail)
        }
    }
}
