package queue

import (
    "../linked-list"
)

type MyDeq interface {
    AddFront(el interface{}) bool
    AddBack(el interface{}) bool
    RemoveFront() interface{}
    RemoveBack() interface{}
    PeekFront() interface{}
    PeekBack() interface{}
    Size() int
}

type Deq struct {
    list *linkedList.LinkedList
}

func (d *Deq) AddFront(el interface{}) bool {
    return d.list.InsertBefore(0, el)
}

func (d *Deq) AddBack(el interface{}) bool {
    return d.list.Add(el)
}

func (d *Deq) RemoveFront() interface{} {
    return d.list.Remove(0)
}

func (d *Deq) RemoveBack() interface{} {
    return d.list.Remove(d.list.Size()-1)
}

func (d *Deq) PeekFront() interface{} {
    return d.list.GetNth(0).Data
}

func (d *Deq) PeekBack() interface{} {
    var back int
    currSize := d.list.Size()
    if currSize == 0 {
        back = 0
    } else {
        back = currSize - 1
    }

    return d.list.GetNth(back).Data
}

func (d *Deq) Size() int {
    return d.list.Size()
}

func (d *Deq) FillList(size int) {
    d.list.FillList(size)
}

func NewDeq() *Deq {
    return  &Deq{ linkedList.NewLinkedList() }
}
