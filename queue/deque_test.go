package queue

import (
    "testing"
)

func TestAddFront(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeq := NewDeq()
        currVal := iterations*2
        if newDeq.AddFront(currVal) {
            curr := newDeq.PeekFront()
            
            if curr != currVal {
               t.Errorf("Expected %v, but got %v", currVal, curr) 
            }
        }
    }
}

func TestAddBack(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeq := NewDeq()
        currVal := iterations*2
        if newDeq.AddBack(currVal) {
            curr := newDeq.PeekBack()
            
            if curr != currVal {
               t.Errorf("Expected %v, but got %v", currVal, curr) 
            }
        }
    }
}

func filledDeq(delta int) *Deq {
    newDeq := NewDeq()
    loops := delta+50
    for loops != 0 {
        newDeq.AddBack(loops)
        loops--
    }

    return newDeq
}

func TestRemoveFront(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeq := filledDeq(iterations)
        oldTop := newDeq.PeekFront()
        removed := newDeq.RemoveFront()
        newVal := newDeq.PeekFront()

        if removed != oldTop || oldTop == newVal {
            t.Errorf("Did not remove %v oldTop%p %+v, newVal%p %+v", removed, oldTop, oldTop, newVal, newVal)
        }
    }
}



func TestRemoveBack(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeq := filledDeq(iterations)
        oldTop := newDeq.PeekBack()
        removed := newDeq.RemoveBack()
        newVal := newDeq.PeekBack()

        if removed != oldTop || oldTop == newVal {
            t.Errorf("Did not remove %v oldTop%p %+v, newVal%p %+v", removed, oldTop, oldTop, newVal, newVal)
        }
    }
}

func TestPeekFront(t *testing.T) {
    newDeq := NewDeq()
    for iterations := 1; iterations < 100; iterations++ {
        newDeq.AddFront(iterations)
        currTop := newDeq.PeekFront()
        if currTop !=  iterations {
            t.Errorf("Expected %+v, but got %v", iterations, currTop)
        }
    }
}

func TestPeekBack(t *testing.T) {
    newDeq := NewDeq()
    for iterations := 1; iterations < 100; iterations++ {
        newDeq.AddBack(iterations)
        currTail := newDeq.PeekBack()
        if currTail !=  iterations {
            t.Errorf("Expected %+v, but got %v", iterations, currTail)
        }
    }
}
