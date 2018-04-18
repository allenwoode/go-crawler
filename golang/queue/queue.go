package queue

type Queue []interface{}

// Push element into queue
// 		e.g. Push(123)
func (q *Queue) Push(v int)  {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return -1
	}

	last := len(*q) - 1
	tail := (*q)[last]
	*q = (*q)[:last]
	return tail.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}