package linkedList

import (
    "testing"
    "math/rand"
    "time"
    //"fmt"
)

func makeLinkedList(withCapacity int, fill bool) (*LinkedList, int) {
    var capacity int
    i := 1
    if withCapacity > 0 {
        capacity = withCapacity
    } else {
        capacity = getRandIntn(0)
    }

    newList := &LinkedList{ nil, nil, 0 }
    if fill {
        for capacity > 0 {
            // what happens if past capacity ??
            if newList.Add(i) {
                i++
            }
            capacity--
        }
    }

    return newList, i-1
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
    var prev, currItem *Node
    var validData, validPrev bool
    for iterations := 0; iterations < 100; iterations++ {
        newList, _ := makeLinkedList(0, true)
        for i := 0; i < 30; i++ {
            if newList.Add(i+1) {
                currItem = newList.GetNth(newList.Size()-1)
                prev = newList.GetNth(newList.Size()-2)

                if currItem != nil {
                    validData = currItem.data == i+1
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

        randIdx := getRandIntn(actualSize)
        currItem := newList.GetNth(randIdx)
        target := newList.head
        for target != nil && randIdx > 0 && target.next != nil {
            target = target.next
            randIdx--
        }

        if currItem != target {
            t.Errorf("Expected %+v, but got %+v\n", target, currItem)
        }
    }
}

//func TestAddAtI(t *testing.T) {
//    var addSuccess bool
//    var prev, next *Node
//    for iterations := 0; iterations < 100; iterations++ {
//        newList, _ := makeLinkedList(0, true)
//        randIdx := getRandIntn(newList.Size())
//        if randIdx > 0 {
//            prev = newList.GetNth(randIdx-1)
//        }
//        
//        if randIdx <= newList.Size()-1 {
//            next = newList.GetNth(randIdx+2)
//        }
//
//        if !newList.AddAtI(randIdx, 50*randIdx) {
//            continue
//        }
//
//        currItem := newList.GetNth(randIdx)
//        fmt.Printf("adding item %+v at %p index %d\n", currItem, currItem, randIdx)
//        fmt.Printf("previous %p ========= %+v\n", prev, prev)
//        fmt.Printf("=========%p %+v\n", currItem, currItem)
//
//        // if currItem not nil, with data, with correct previous and next
//        addSuccess = currItem != (&Node{}) && currItem.data == 50*randIdx &&
//        currItem.prev == prev && currItem.next == next
//        
//        if prev != nil {
//            addSuccess = currItem == prev.next
//        }
//
//        if next != nil {
//            addSuccess = currItem == next.prev
//        }
//
//        fmt.Println("---------------------------------")
//        //newList.PrintList()
//        if !addSuccess {
//            t.Errorf("tried idx %d with list size %d Prev %p: %+v\nCurr %p: %+v\nNext %p: %+v", randIdx, newList.Size(), prev, prev, currItem, currItem, next, next)
//        }
//        fmt.Println("00000000000000000000000000")
//    }
//}
