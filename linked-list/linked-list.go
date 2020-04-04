package linkedList

import "fmt"

type MyLinkedList interface {
    Size() int
    GetNth(i int) *Node
    Add(el interface{}) bool
    AddAtI(i int, el interface{}) bool
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

// TODO TEST THIS
func IsNil(el interface{}) bool {
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

    return nilVal == el
}

// returns pointer to the value of Node
func (l *LinkedList) GetNth(i int) *Node {
    if i < 0 || i >= l.Size() {
        return nil
    }

    target := l.head;
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
        
    if l.head == nil {
        l.head = &Node{nil, nil, el}
    } else {
        if l.tail == nil {
            temp = l.head
        } else {
            temp = l.tail
        }

        l.tail = &Node{temp, nil, el}
        l.tail.prev = temp
        temp.next = l.tail
    }

    l.size++
    return true
}

// @params i int: index to insert el
func (l *LinkedList) AddAtI(i int, el interface{}) bool {
    if IsNil(el) {
        return false
    }

    // handle index out of range
    if i < 0 || i > l.Size() {
        return false
    }

    curr := l.GetNth(i)
    prev := curr.prev
    next := curr
    newNode := &Node{prev, curr, el}
        fmt.Printf("curr====%+v\n", curr)
        fmt.Printf("newnode====%+v\n", newNode)
        fmt.Printf("prev====%+v\n", prev)
    if prev != nil {
        prev.next = newNode
    }

    if next != nil {
        next.prev = newNode
    }

    if i == 0 {
        l.head = newNode
    } else if l.tail == nil {
        l.tail = newNode
    }

    l.size++
    return true
}

func (l *LinkedList) PrintList() {
    target := l.head;
    fmt.Printf("head==== %p %+v\n", target, target)
    target = target.next
    for target != nil {
        fmt.Printf("==== %p %+v\n", target, target)
        target = target.next
    }
    fmt.Printf("tail==== %p %+v\n", l.tail, l.tail)
    //fmt.Printf("==== %p %+v\n", target, target)
    fmt.Println("======================== -- =============================")
}
