package homework

import (
	"fmt"
	"testing"
)

func Test_MatrixArray(t *testing.T) {
	m := NewMatrixArray[int](5)
	_ = m.Enqueue(0, 10)
	_ = m.Enqueue(1, 2)
	_ = m.Enqueue(1, 3)
	fmt.Println(m)
	res, _ := m.Dequeue()
	fmt.Println(res)

	res, _ = m.Dequeue()
	fmt.Println(res)

	res, _ = m.Dequeue()
	fmt.Println(res)

	res, err := m.Dequeue()
	if err != nil {
		fmt.Println(err)
	}

}
