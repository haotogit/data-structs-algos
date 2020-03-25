package main

type LinkedListMethods interface {
    Size() int
    Get(i int) *Node
    Add()
    AddAtI(i int, el Node)
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

    // TODO add(el)
    // append el to list

    // TODO add(i, el)
    // append el to ith position in list

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
    data int
}

type LinkedList struct {
    list []Node
}

func main() {
    // a linked list is a sequence of nodes
    // that contain references to one another
    // double || single, with references to prev && next || next

    // unlike an array, linked lists do not provide constant time access to
    // a specific index because you have to traverse through K items

    // benefit of linkedlist is add and remove from beginning of list is constant

    // type Node
    // type linkedList
}

func (l *LinkedList) Size() int {
    return len(l.list)
}

func (l *LinkedList) Get(i int) Node {
    return l.list[i]
}

func (l *LinkedList) Add(el Node) {
    // handle null value
    // handle increasing capacity
    l.list = append(l.list, el)
}

func (l *LinkedList) AddAtI(i int, el Node) {
    // handle null value
    // handle increasing capacity
    // handle index out of range
    l.Add(el)
    if i < len(l.list) {
        copy(l.list[i+1:], l.list[i:])
        l.list[i] = el
    }
    //l.Add(el)
}
