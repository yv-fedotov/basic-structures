package homework

import "errors"

type FactoryArray[T any] struct {
	idx   int
	array []T
}

func NewFactoryArray[T any](size int) *FactoryArray[T] {
	return &FactoryArray[T]{
		idx:   0,
		array: make([]T, size),
	}
}

func (f *FactoryArray[T]) Add(element T) {
	if cap(f.array) == f.idx {
		{
			newSize := len(f.array) * 2
			newArray := make([]T, newSize)
			copy(newArray, f.array)
			f.array = newArray
		}
	}

	f.array[f.idx] = element
	f.idx++
}

func (f *FactoryArray[T]) Remove() (T, error) {
	var result T
	if len(f.array) == 0 {
		return result, errors.New("array is empty")
	}

	f.idx--
	result = f.array[f.idx]

	f.array = f.array[:f.idx]
	return result, nil
}
