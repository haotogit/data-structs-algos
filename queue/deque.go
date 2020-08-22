package queue

import (
    "github.com/haotogit/data-structs-algos/util"
)

type MyDeq interface {
    AddFront(el interface{}) bool
    AddBack(el interface{}) bool
    RemoveFront() interface{}
    RemoveBack() interface{}
    PeekFront() interface{}
    PeekBack() interface{}
    Size() int
    Capacity() int
}

type Node struct {
    prev *Node
    next *Node
    Data interface{}
}

type Deq struct {
    list []*Node
    currSize int
    capacity int
}

func newNode(data interface{}) *Node {
    return &Node{ nil, nil, data }
}

func (d *Deq) AddFront(el interface{}) bool {
    if util.IsNil(el) {
        return false
    }

    if d.currSize == d.capacity {
        // +1 for if cap == 0
        d.capacity = (d.capacity+1)*2
        d.list = make([]*Node, 0, d.capacity)
    }
    
    d.currSize++
    d.list = d.list[:d.currSize]
    if d.currSize != 0 {
        copy(d.list[1:], d.list)
    }

    d.list[0] = newNode(el)
    return true 
}

func (d *Deq) AddBack(el interface{}) bool {
    if util.IsNil(el) {
        return false
    }

    if d.currSize == d.capacity {
        d.capacity = (d.capacity+1)*2
        d.list = make([]*Node, 0, d.capacity)
    }
    
    d.currSize++
    d.list = d.list[:d.currSize]
    d.list[d.currSize-1] = newNode(el)
    return true 
}

func (d *Deq) RemoveFront() interface{} {
    if d.currSize == 0 {
        return nil 
    }

    toRemove := d.list[0]
    // to append or to copy ?
    //d.list = append(d.list[1:], nil)
    copy(d.list[:], d.list[1:])
    d.list[d.currSize-1] = nil
    d.currSize--
    return toRemove.Data
}

func (d *Deq) RemoveBack() interface{} {
    if d.currSize == 0 {
        return nil 
    }

    toRemove := d.list[d.currSize-1]
    d.list[d.currSize-1] = nil
    d.currSize--
    return toRemove.Data
}

func (d *Deq) PeekFront() interface{} {
    if d.currSize == 0 {
        return nil
    }

    return d.list[0].Data
}

func (d *Deq) PeekBack() interface{} {
    if d.currSize == 0 || d.list[d.currSize-1] == nil {
        return nil
    }

    return d.list[d.currSize-1].Data
}

func (d *Deq) Size() int {
    return d.currSize
}

func (d *Deq) Capacity() int {
    return d.capacity
}

func NewDeq(capacity int) *Deq {
    return  &Deq{ make([]*Node, capacity), 0, capacity }
}
