package queue

type Queue []int

func (q *Queue) Push(v int)  {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	if q.IsEmpty() {
		return -1
	}

	last := len(*q) - 1
	tail := (*q)[last]
	*q = (*q)[:last]
	return tail
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}