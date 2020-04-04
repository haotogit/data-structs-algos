package queue

type MyQueue interface {
    MyDeque
    Enqueue(el interface{}) bool
    Dequeue() interface{}
    Peek() interface{}
}

type Queue struct {
    Deque
}

func (q *Queue) Enqueue(el interface{}) bool {
    return q.AddBack(el)
}

func (q *Queue) Dequeue() interface{} {
    return q.RemoveFront()
}

func (q *Queue) Peek() interface{} {
    return q.PeekFront()
}
