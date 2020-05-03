package queue

type MyQueue interface {
    Enqueue(el interface{}) bool
    Dequeue() interface{}
    Peek() interface{}
    Size() int
    Capacity() int
}

type Queue struct {
    items *Deq
}

func (q *Queue) Enqueue(el interface{}) bool {
    return q.items.AddBack(el)
}

func (q *Queue) Dequeue() interface{} {
    return q.items.RemoveFront()
}

func (q *Queue) Peek() interface{} {
    return q.items.PeekFront()
}

func (q *Queue) Size() int {
    return q.items.Size()
}

func (q *Queue) Capacity() int {
    return q.items.Capacity()
}

// if capacity bounded queue, and using 
// slice as backing store, but if no capacity
// wouldnt be bad to use DeqLL since actions
// are only on head and tail of list
func NewQ(capacity int) *Queue {
    return &Queue{ NewDeq(capacity) }
}
