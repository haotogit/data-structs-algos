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
}

func (d *DeqLL) AddFront(el interface{}) bool {
    if util.IsNil(el) {
        return false
    }

    if d.list.InsertBefore(0, el) {
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
        return true
    } else {
        return false
    }
}

func (d *DeqLL) RemoveFront() interface{} {
    return d.list.Remove(0)
}

func (d *DeqLL) RemoveBack() interface{} {
    return d.list.Remove(d.Size()-1)
}

func (d *DeqLL) PeekFront() interface{} {
    if d.Size() == 0 {
        return nil
    }

    return d.list.GetNth(0).Data
}

func (d *DeqLL) PeekBack() interface{} {
    if d.Size() == 0 {
        return nil
    }

    return d.list.GetNth(d.Size()-1).Data
}

func (d *DeqLL) Size() int {
    return d.list.Size()
}

func (d *DeqLL) FillList(size int) {
    d.list.FillList(size)
}

func NewDeqLL() *DeqLL {
    return  &DeqLL{ linkedList.NewLinkedList() }
}
