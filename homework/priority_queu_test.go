package homework

import (
	"fmt"
	"testing"
)

func Test_PriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()

	q.Enqueue(1, 1)
	q.Enqueue(2, 2)
	q.Enqueue(2, 3)

	fmt.Println(q.root.value)

	e := q.Dequeue()
	fmt.Println(e)

	e = q.Dequeue()
	fmt.Println(e)

}