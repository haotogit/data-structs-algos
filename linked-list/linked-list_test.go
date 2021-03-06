package linkedList

import (
    "testing"
    "github.com/haotogit/data-structs-algos/util"
)

func makeLinkedList(withSize int, fill bool) (*LinkedList, int) {
    var currSize int
    i := 1
    if withSize > 0 {
        currSize = withSize
    } else {
        currSize = util.GetRandIntn(0, 100) + 25
    }

    newList := NewLinkedList()
    if fill {
        for currSize > 0 {
            // what happens if past capacity ??
            if newList.Add(i) {
                i++
            }
            currSize--
        }
    }

    return newList, i-1
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
    var prev, currItem *Node
    var validData, validPrev bool
    for iterations := 0; iterations < 100; iterations++ {
        newList, _ := makeLinkedList(0, true)
        for i := 0; i < 30; i++ {
            if newList.Add(i+1) {
                currItem = newList.GetNth(newList.Size()-1)
                prev = newList.GetNth(newList.Size()-2)

                if currItem != nil {
                    validData = currItem.Data == i+1
                    validPrev = currItem.prev == prev
                }

                if !validData || (prev != nil && !validPrev) {
                    t.Errorf("prev(%p) %+v, newNode(%p): %+v\n", prev, prev, currItem, currItem)
                }
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

        randIdx := util.GetRandIntn(0, actualSize)
        currItem := newList.GetNth(randIdx)
        target := newList.Head
        for target != nil && randIdx > 0 && target.next != nil {
            target = target.next
            randIdx--
        }

        if currItem != target {
            t.Errorf("Expected %+v, but got %+v\n", target, currItem)
        }
    }
}

func TestInsertBefore(t *testing.T) {
    var addSuccess bool
    var prev, next *Node
    for iterations := 0; iterations < 100; iterations++ {
        newList, _ := makeLinkedList(0, true)
        randIdx := util.GetRandIntn(0, newList.Size())

        if newList.InsertBefore(randIdx, 50*randIdx) {
            if randIdx > 0 {
                prev = newList.GetNth(randIdx-1)
            }

            if randIdx <= newList.Size()-2 {
                next = newList.GetNth(randIdx+1)
            }

            currItem := newList.GetNth(randIdx)
            // if currItem not nil, with data, with correct previous and next
            addSuccess = currItem != nil && currItem.Data == 50*randIdx &&
                currItem.prev == prev && currItem.next == next

            if addSuccess {
                if prev != nil {
                    addSuccess = currItem == prev.next
                }

                if next != nil {
                    addSuccess = currItem == next.prev
                }
            }

            if !addSuccess {
                t.Errorf("tried idx %d with list size %d Prev %p: %+v\nCurr %p: %+v\nNext %p: %+v", randIdx, newList.Size(), prev, prev, currItem, currItem, next, next)
            }
        }
    }
}

func getDerivedVal(el interface{}) interface{} {
    switch el.(type) {
        case string:
            return el.(string)+"changessssss"
        case int:
            return el.(int)*3
        default:
            // had to include this because of t...
            return nil
    }
}

func TestSet(t *testing.T) {
    for iterations := 1; iterations <= 100; iterations++ {
        newList, _ := makeLinkedList(0, true)
        randIdx := util.GetRandIntn(0, newList.Size()-1)
        currNode := newList.GetNth(randIdx)
        newVal := getDerivedVal(currNode.Data)
        oldVal := newList.Set(randIdx, newVal)

        if oldVal == currNode.Data || currNode.Data != newVal {
            t.Errorf("Did not set %v properly oldVal %v, currNode%+v", newVal, oldVal, currNode)
        }
    }
}

func TestRemove(t *testing.T) {
    for iterations := 1; iterations <= 100; iterations++ {
        newList, _ := makeLinkedList(0, true)
        oldSize := newList.Size()
        for removals := 0; removals <= newList.Size()/2; removals++ {
            randIdx := util.GetRandIntn(0, oldSize-1)
            currNode := newList.GetNth(randIdx)

            removedItem := newList.Remove(randIdx)
            newVal := newList.GetNth(randIdx)

            if removedItem != currNode.Data || currNode.Data == newVal || newList.Size() != oldSize-1 {
                t.Errorf("Failed to remove at index %d, got %v, newVal %v", randIdx, removedItem, newVal)
            }

            oldSize--
        }
    }
}
