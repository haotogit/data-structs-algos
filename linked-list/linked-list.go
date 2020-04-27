package linkedList

import (
    "fmt"
    "../util"
)

type MyLinkedList interface {
    Size() int
    GetNth(i int) *Node
    Add(el interface{}) bool
    InsertBefore(i int, el interface{}) bool
    Set(i int, el interface{}) interface{}
    Remove(i int) interface{}
}

type Node struct {
    prev *Node
    next *Node
    Data interface{}
}

type LinkedList struct {
    Head *Node
    Tail *Node
    size int
}

func (l *LinkedList) Size() int {
    return l.size 
}

func IsNil(el interface{}) bool {
    var nilVal interface{}
    switch el.(type) {
        case string:
            nilVal = ""
            // had to include this because of t...
        case int:
            nilVal = 0
        default:
            nilVal = nil
    }

    return nilVal == el
}

// returns pointer to the value of Node
func (l *LinkedList) GetNth(i int) *Node {
    if i < 0 || i >= l.Size() {
        return nil
    }

    target := l.Head;
    for i > 0 && target != nil && target.next != nil {
        target = target.next
        i--
    }

    return target
}

func (l *LinkedList) Add(el interface{}) bool {
    var temp *Node
    if IsNil(el) {
        return false
    }
        
    if l.Head == nil {
        l.Head = &Node{nil, nil, el}
    } else {
        if l.Tail == nil {
            temp = l.Head
        } else {
            temp = l.Tail
        }

        l.Tail = &Node{temp, nil, el}
        l.Tail.prev = temp
        temp.next = l.Tail
    }

    l.size++
    return true
}

// @params i int: index to insert el
func (l *LinkedList) InsertBefore(i int, el interface{}) bool {
    var newNode, prev *Node
    if IsNil(el) {
        return false
    }

    // handle index out of range
    if i < 0 || i > l.Size() {
        return false
    }

    curr := l.GetNth(i)
    if curr != nil {
        prev = curr.prev
        newNode = &Node{prev, curr, el}
        curr.prev = newNode

        if prev != nil {
            prev.next = newNode
        }

        if i == 0 {
            l.Head = newNode
        }

        if i == l.Size()-1 {
            l.Tail = curr
        }
    } else if i == l.Size() {
        //adding at the end of list
        if i == 0 {
            newNode = &Node{nil, nil, el}
            l.Head = newNode
        } else {
            if l.Tail != nil {
                prev = l.Tail
            } else {
                prev = l.Head
            }
            newNode = &Node{prev, nil, el}
            prev.next = newNode
            l.Tail = newNode
        }
    }

    l.size++
    return true
}

func (l *LinkedList) Set(i int, el interface{}) interface{} {
    var oldVal interface{}
    // do i need to panic instead ??
    if IsNil(el) {
        return nil
    }

    currNode := l.GetNth(i)
    if currNode == nil {
        return nil
    }

    oldVal = currNode.Data
    currNode.Data = el
    return oldVal
}

func (l *LinkedList) Remove(i int) interface{} {
    // do i need to panic instead ??
    currNode := l.GetNth(i)
    if currNode == nil {
        return nil
    }

    prev := l.GetNth(i-1)
    next := l.GetNth(i+1)

    if prev != nil {
        prev.next = next

        if i == l.Size()-1 {
            l.Tail = prev
        }
    }
    
    if next != nil {
        next.prev = prev

        if i == 0 {
            l.Head = next
        }
    }

    l.size--
    return currNode.Data
}

func (l *LinkedList) FillList(size int) {
    var withSize int
    if size > 0 {
        withSize = size 
    } else {
        withSize = util.GetRandIntn(100) + 25
    }

    for withSize > 0 {
        l.Add(util.GetRandIntn(1000)+15)
        withSize--
    }
}

func (l *LinkedList) PrintList() {
    target := l.Head;
    fmt.Printf("head==== %p %+v\n", target, target)
    target = target.next
    for target != nil {
        fmt.Printf("==== %p %+v\n", target, target)
        target = target.next
    }
    fmt.Printf("tail==== %p %+v\n", l.Tail, l.Tail)
    //fmt.Printf("==== %p %+v\n", target, target)
    fmt.Println("======================== -- =============================")
}

func NewLinkedList() *LinkedList {
    return &LinkedList{ nil, nil, 0 }
} 
