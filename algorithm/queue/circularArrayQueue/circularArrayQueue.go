package circularArrayQueue

type CircularArrayQueue struct {
	array []interface{}
	n     int
	head  int
	tail  int
}

func NewCircularArrayQueue(cap int) *CircularArrayQueue {
	return &CircularArrayQueue{
		array: make([]interface{}, cap),
		n:     cap,
		head:  0,
		tail:  0,
	}
}

//入队
func (q *CircularArrayQueue) Enqueue(item interface{}) bool {
	//队列已满
	if (q.tail+1)%q.n == q.head {
		return false
	}
	q.array[q.tail] = item
	q.tail = (q.tail + 1) % q.n
	return true
}

//出队
func (q *CircularArrayQueue) Dequeue(item interface{}) interface{} {
	//队列已空
	if q.head == q.tail {
		return nil
	}
	ret := q.array[q.head]
	q.head = (q.head + 1) % q.n
	return ret
}
