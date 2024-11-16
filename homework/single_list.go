package homework

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkList[T any] struct {
	head *Node[T]
}

func NewLinkList[T any]() *LinkList[T] {
	return &LinkList[T]{}
}

func (l *LinkList[T]) Push(element T) {
	node := &Node[T]{value: element}
	if l.head == nil {
		l.head = node
		return
	}

	p := l.head
	for p.next != nil {
		p = p.next
	}
	p.next = node
}

func (l *LinkList[T]) Pop() T {
	var result T
	if l.head == nil {
		return result
	}

	prev := &Node[T]{}

	for l.head.next != nil {
		prev = l.head
		l.head = l.head.next
	}

	result = l.head.value
	prev.next = nil
	l.head = prev
	return result
}
