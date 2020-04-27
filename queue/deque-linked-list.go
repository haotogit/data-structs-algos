package queue

// deque with linkedlist is not ideal because
// when add, rm, or get in not head or tail of list,
// it's an O(currIndex) opposed to indexed array
import (
    "../linked-list"
    "../util"
)

type MyDeqLL interface {
    AddFront(el interface{}) bool
    AddBack(el interface{}) bool
    RemoveFront() interface{}
    RemoveBack() interface{}
    PeekFront() interface{}
    PeekBack() interface{}
    Size() int
}

type DeqLL struct {
    list *linkedList.LinkedList
    currSize int
    capacity int
}

func (d *DeqLL) AddFront(el interface{}) bool {
    if util.IsNil(el) {
        return false
    }

    if d.list.InsertBefore(0, el) {
        d.currSize++
        return true
    } else {
        return false
    }
}

func (d *DeqLL) AddBack(el interface{}) bool {
    if util.IsNil(el) {
        return false
    }

    if d.list.Add(el) {
        d.currSize++
        return true
    } else {
        return false
    }
}

func (d *DeqLL) RemoveFront() interface{} {
    if d.currSize == 0 {
        return nil 
    }

    removed := d.list.Remove(0)
    if removed != nil {
        d.currSize--
    }

    return removed
}

func (d *DeqLL) RemoveBack() interface{} {
    if d.currSize == 0 {
        return nil
    }

    removed := d.list.Remove(d.currSize-1)
    if removed != nil {
        d.currSize--
    }

    return removed
}

func (d *DeqLL) PeekFront() interface{} {
    if d.currSize == 0 {
        return nil
    }

    return d.list.GetNth(0).Data
}

func (d *DeqLL) PeekBack() interface{} {
    if d.currSize == 0 {
        return nil
    }

    return d.list.GetNth(d.currSize-1).Data
}

func (d *DeqLL) Size() int {
    return d.currSize
}

func (d *DeqLL) FillList(size int) {
    d.list.FillList(size)
}

func NewDeqLL(capacity int) *DeqLL {
    return  &DeqLL{ linkedList.NewLinkedList(), 0, capacity }
}
