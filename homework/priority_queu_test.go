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
	q.Enqueue(1, 4)

	for i := 0; i < 4; i++ {
		e := q.Dequeue()
		fmt.Println(e)
	}

}
