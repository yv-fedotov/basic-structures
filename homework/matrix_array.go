package homework

import "errors"

type MatrixArray[T any] struct {
	array [][]T
	size  int
}

func NewMatrixArray[T any](size int) *MatrixArray[T] {
	array := make([][]T, size)

	return &MatrixArray[T]{
		array: array,
		size:  size,
	}
}

func (m *MatrixArray[T]) Enqueue(priority int, element T) error {
	if priority < 0 || priority >= m.size {
		return errors.New("priority out of range")
	}

	if len(m.array[priority]) == 0 {
		m.array[priority] = []T{element}
	} else {
		item := m.array[priority]
		item = append(item, element)
		m.array[priority] = item
	}

	return nil
}

func (m *MatrixArray[T]) Dequeue() (T, error) {
	var result T
	if len(m.array) == 0 {
		return result, errors.New("empty array")
	}
	for i := range m.array {
		if len(m.array[i]) == 0 {
			continue
		}
		for j := range m.array[i] {
			result = m.array[i][j]
			item := m.array[i]
			item = append(item[:j], item[j+1:]...)
			m.array[i] = item
			return result, nil
		}
	}
	return result, errors.New("empty array")
}
