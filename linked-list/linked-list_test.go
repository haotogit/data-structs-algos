package linkedList

import (
    "testing"
    "math/rand"
    "time"
)

func makeLinkedList(withCapacity int, fill bool) (*LinkedList, int) {
    var capacity int
    i := 0
    if withCapacity > 0 {
        capacity = withCapacity
    } else {
        capacity = getRandIntn(0)
    }

    newList := &LinkedList{ nil, nil, 0 }
    if fill {
        for i < capacity {
            // what happens if past capacity ??
            newList.Add(Node{nil, nil, i})
            i++
        }
    }

    return newList, i
}

func getRandIntn(max int) int {
    rand.Seed(time.Now().UnixNano())
    if max == 0 {
        return rand.Intn(50)
    } else {
        return rand.Intn(max)
    }
}

func TestSize(t *testing.T) {
    for iterations := 0; iterations < 100; iterations++ {
        newList, actualSize := makeLinkedList(0, true)
        if newList.Size() != actualSize {
            t.Errorf("Expected: %d Actual: %d\n", actualSize, newList.Size())
        }
    }
}

func TestAdd(t *testing.T) {
    for iterations := 0; iterations < 100; iterations++ {
        newList, _ := makeLinkedList(0, false)
        for i := 1; i < 10; i++ {
            if !newList.Add(i) {
                t.Errorf("Failed to add %d\n", i)
            }
        }
    }
}

func TestAddNil(t *testing.T) {
    newList, _ := makeLinkedList(0, false)
    if newList.Add(0) || newList.Add("") {
        t.Errorf("No nil value allowed\n")
    }
}

func TestGetNth(t *testing.T) {
    for iterations := 0; iterations < 100; iterations++ {
        newList, actualSize := makeLinkedList(0, true)
        if actualSize > 1 {
            actualSize--
        }

        randIdx := getRandIntn(actualSize)
        currItem := newList.GetNth(randIdx)
        target := newList.head
        for target != nil && randIdx >= 0 && target.next != nil {
            target = target.next
            randIdx--
        }

        if currItem != target {
            t.Errorf("Expected %+v, but got %+v\n", target, currItem)
        }
    }
}
