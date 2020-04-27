package queue

type MyQueue interface {
    Enqueue(el interface{}) bool
    Dequeue() interface{}
    Peek() interface{}
    Size() int
    Capacity() int
}

type Queue struct {
    currDeq *Deq
}

func (q *Queue) Enqueue(el interface{}) bool {
    return q.currDeq.AddBack(el)
}

func (q *Queue) Dequeue() interface{} {
    return q.currDeq.RemoveFront()
}

func (q *Queue) Peek() interface{} {
    return q.currDeq.PeekFront()
}

func (q *Queue) Size() int {
    return q.currDeq.Size()
}

func (q *Queue) Capacity() int {
    return q.currDeq.Capacity()
}

func NewQ(capacity int) *Queue {
    return &Queue{ NewDeq(capacity) }
}
