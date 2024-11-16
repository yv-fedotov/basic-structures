package homework

import "fmt"

type SingleArray[T any] struct {
	array []T
}

func NewSingleArray[T any](size int) *SingleArray[T] {
	return &SingleArray[T]{
		array: make([]T, size),
	}
}

func (s *SingleArray[T]) Add(element T, index int) error {
	if index >= len(s.array) {
		return fmt.Errorf("index out of range")
	}
	s.array[index] = element
	return nil
}

func (s *SingleArray[T]) Remove(index int) (T, error) {
	var result T
	if index >= len(s.array) {

		return result, fmt.Errorf("index out of range")
	}
	if index < 0 {
		return result, fmt.Errorf("array is empty")
	}
	element := s.array[index]
	s.array[index] = result
	return element, nil
}
