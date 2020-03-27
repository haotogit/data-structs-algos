package linkedList

import "fmt"

type MyLinkedList interface {
    Size() int
    GetNth(i int) *Node
    Add(el interface{}) bool
    AddAtI(i int, el interface{})
    //Set(i int, el)
    //Remove(i int, el)
    //Clear()

    // TODO getNth(i)
    // params
    // i - index for desired element
    // return nth node

    // TODO get(i)
    // i - index for desired element
    // return element

    // TODO set(i, el)
    // set el to ith position in list

    // TODO remove(i, el)
    // set el to ith position in list

    // TODO clear()
    // remove all items from list

    // TODO iterator
}

type Node struct {
    prev *Node
    next *Node
    data interface{}
}

type LinkedList struct {
    head *Node
    tail *Node
    size int
}

func (l *LinkedList) Size() int {
    return l.size 
}

func isNil(el interface{}) interface{} {
    var nilVal interface{}
    switch t := el.(type) {
        case string:
            nilVal = ""
            // had to include this because of t...
            fmt.Println(t)
        case int:
            nilVal = 0
        default:
            nilVal = nil
    }

    return nilVal
}

func (l *LinkedList) GetNth(i int) *Node {
    target := l.head;
    for target != nil && i >= 0 && target.next != nil {
        target = target.next
        i--
    }

    return target
}

func (l *LinkedList) Add(el interface{}) bool {
    if el == isNil(el) {
        return false
    }

    if l.head == nil {
        l.head = &Node{nil, nil, el}
    } else if l.tail == nil {
        // why not &l.head
        l.tail = &Node{l.head, nil, el}
    } else {
        target := &Node{l.tail, nil, el}
        l.tail.next = target
    }

    l.size++
    return true
}

//func (l *LinkedList) AddAtI(i int, el Node) {
//    // handle null value
//    // handle increasing capacity
//    // handle index out of range
//    l.Add(el)
//    if i < len(l.list) {
//        copy(l.list[i+1:], l.list[i:])
//        l.list[i] = el
//    }
//    //l.Add(el)
//}
