package queue

// Queue
type Queue []interface{}

// Push an element into queue
// 		e.g. Push(123)
func (q *Queue) Push(v int)  {
	*q = append(*q, v)
}

// Pop an element out queue
//		e.g. Pop()
func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return -1
	}
	last := len(*q) - 1
	tail := (*q)[last]
	*q = (*q)[:last]
	return tail.(int)
}

// Check queue is not empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}