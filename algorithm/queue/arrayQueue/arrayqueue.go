package arrayQueue

//用数组实现一个顺序栈
type ArrayQueue struct {
	array []interface{}
	cap   int
	head  int
	tail  int
}

func NewArrayQueue(cap int) *ArrayQueue {
	return &ArrayQueue{
		array: make([]interface{}, cap),
		cap:   cap,
		head:  0,
		tail:  0,
	}
}

func (q *ArrayQueue) Enqueue(item interface{}) bool {
	// 如果tail == n 表示队列已经到了队尾
	if q.tail == q.cap {
		// 如果head在0处，表示队列已经满了
		if q.head == 0 {
			return false
		}
		for i := q.head; i < q.tail; i++ {
			q.array[i-q.head] = q.array[i]
		}
		q.tail = q.tail - q.head
		q.head = 0
	}
	q.array[q.tail] = item
	q.tail++
	return true
}

func (q *ArrayQueue) Dequeue() interface{} {
	// 如果head == tail 表示队列为空
	if q.head == q.tail {
		return nil
	}
	ret := q.array[q.head]
	q.head++
	return ret
}
