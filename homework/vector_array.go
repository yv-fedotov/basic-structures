package homework

import "fmt"

type VectorArray[T any] struct {
	array []T
}

func NewVectorArray[T any](size int) *VectorArray[T] {
	return &VectorArray[T]{
		array: make([]T, 0, size),
	}
}

func (v *VectorArray[T]) Add(element T) {
	v.array = append(v.array, element)
}

func (v *VectorArray[T]) Remove(index int) (T, error) {
	var result T
	if len(v.array) == 0 {
		return result, fmt.Errorf("array is empty")
	}
	if len(v.array) <= index {
		return result, fmt.Errorf("index out of range")
	}
	element := v.array[index]
	v.array = append(v.array[:index], v.array[index+1:]...)
	return element, nil
}
