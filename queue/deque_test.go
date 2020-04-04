package queue

import (
    "testing"
)

func TestAddFront(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeque := &Deque{}
        currVal := iterations*2
        if newDeque.AddFront(currVal) {
            curr := newDeque.GetNth(0)
            
            if curr.Data != currVal {
               t.Errorf("Expected %v, but got %v", currVal, curr.Data) 
            }
        }
    }
}

func TestAddBack(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeque := &Deque{}
        currVal := iterations*2
        if newDeque.AddBack(currVal) {
            curr := newDeque.GetNth(newDeque.Size()-1)
            
            if curr.Data != currVal {
               t.Errorf("Expected %v, but got %v", currVal, curr.Data) 
            }
        }
    }
}

func filledDeque(delta int) *Deque {
    newDeque := &Deque{}
    loops := delta+50
    for loops != 0 {
        newDeque.AddBack(loops)
        loops--
    }

    return newDeque
}

func TestRemoveFront(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeque := filledDeque(iterations)
        oldTop := newDeque.GetNth(0)
        removed := newDeque.RemoveFront()
        newVal := newDeque.GetNth(0)

        if removed != oldTop.Data || oldTop == newVal {
            t.Errorf("Did not remove %v oldTop%p %+v, newVal%p %+v", removed, oldTop, oldTop, newVal, newVal)
        }
    }
}



func TestRemoveBack(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeque := filledDeque(iterations)
        oldTop := newDeque.GetNth(newDeque.Size()-1)
        removed := newDeque.RemoveBack()
        newVal := newDeque.GetNth(newDeque.Size()-1)

        if removed != oldTop.Data || oldTop == newVal {
            t.Errorf("Did not remove %v oldTop%p %+v, newVal%p %+v", removed, oldTop, oldTop, newVal, newVal)
        }
    }
}

func TestPeekFront(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeque := filledDeque(iterations)
        currTop := newDeque.PeekFront()
        expected := newDeque.GetNth(0)
        if currTop !=  expected.Data {
            t.Errorf("Expected %+v, but got %v", expected, currTop)
        }
    }
}

func TestPeekBack(t *testing.T) {
    for iterations := 1; iterations < 100; iterations++ {
        newDeque := filledDeque(iterations)
        currTail := newDeque.PeekBack()
        expected := newDeque.GetNth(newDeque.Size()-1)
        if currTail !=  expected.Data {
            t.Errorf("Expected %+v, but got %v", expected, currTail)
        }
    }
}
