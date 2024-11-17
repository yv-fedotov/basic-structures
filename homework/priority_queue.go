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

	pointer := q.root

	for pointer.next != nil {
		if pointer.priority >= priority {
			node.next = pointer.next
			pointer.next = node
			return
		}

		pointer = pointer.next
	}

	pointer.next = node
}

func (q *PriorityQueue[T]) Dequeue() T {
	var result T
	if q.root == nil {
		return result
	}

	result = q.root.value
	q.root = q.root.next

	return result
}
