package main

import (
    "testing"
    "math/rand"
    "time"
)

func makeLinkedList(withCapacity int) (LinkedList, int) {
    var newList LinkedList
    var i, capacity int
    rand.Seed(time.Now().UnixNano())
    if withCapacity > 0 {
        capacity = withCapacity
    } else {
        capacity = rand.Intn(50)
    }

    newList = LinkedList{ make([]Node, 0, capacity) }
    i = 0
    for i < capacity {
        // what happens if past capacity ??
        newList.list = append(newList.list, Node{nil, nil, i})
        i++
    }

    return newList, i
}

func TestSize(t *testing.T) {
    var newList LinkedList
    var actualSize int
    for iterations := 0; iterations < 100; iterations++ {
        newList, actualSize = makeLinkedList(0)
        if newList.Size() != actualSize {
            t.Errorf("Expected: %d Actual: %d\n", actualSize, newList.Size())
        }
    }
}

func TestGet(t *testing.T) {
    var newList LinkedList
    var actualSize, randIdx int
    var currItem Node
    for iterations := 0; iterations < 100; iterations++ {
        newList, actualSize = makeLinkedList(0)
        if actualSize == 0 {
            continue
        } else if actualSize > 1 {
            actualSize--
        }

        rand.Seed(time.Now().UnixNano())
        randIdx = rand.Intn(actualSize)
        currItem = newList.Get(randIdx)
        if currItem != newList.list[randIdx] {
            t.Errorf("Expected %p but got %p\n", &newList.list[randIdx], &currItem)
        }
    }
}
