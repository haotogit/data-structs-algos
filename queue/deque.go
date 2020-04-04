package queue

import (
    "../linked-list"
)

type MyDeque interface {
    linkedList.MyLinkedList
    AddFront(el interface{}) bool
    AddBack(el interface{}) bool
    RemoveFront() interface{}
    RemoveBack() interface{}
    PeekFront() interface{}
    PeekBack() interface{}
}

type Deque struct {
    linkedList.LinkedList
}

func (d *Deque) AddFront(el interface{}) bool {
    return d.InsertBefore(0, el)
}

func (d *Deque) AddBack(el interface{}) bool {
    return d.Add(el)
}

func (d *Deque) RemoveFront() interface{} {
    return d.Remove(0)
}

func (d *Deque) RemoveBack() interface{} {
    return d.Remove(d.Size()-1)
}

func (d *Deque) PeekFront() interface{} {
    if d.Head == nil {
        return nil
    }

    return d.Head.Data
}

func (d *Deque) PeekBack() interface{} {
    if d.Tail == nil {
        return nil
    }

    return d.Tail.Data
}
