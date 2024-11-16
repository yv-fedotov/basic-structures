package homework

type NodeQueue[T any] struct {
	value    T
	priority int
	next     *NodeQueue[T]
}

type PriorityQueue[T any] struct {
	root *NodeQueue[T]
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

func (q *PriorityQueue[T]) Enqueue(priority int, value T) {
	node := &NodeQueue[T]{
		value:    value,
		priority: priority,
	}

	if q.root == nil {
		q.root = node
		return
	}

	node.next = q.root

	q.root = node

}

func (q *PriorityQueue[T]) Dequeue() T {
	var result T
	if q.root == nil {
		return result
	}

	prev := &NodeQueue[T]{}
	head := q.root
	for head.next != nil {
		prev = head
		head = head.next
	}

	result = head.value
	prev.next = nil
	head = prev

	return result
}
