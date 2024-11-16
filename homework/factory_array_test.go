package homework

import (
	"fmt"
	"testing"
)

func Test_FactoryArray(t *testing.T) {
	f := NewFactoryArray[int](3)
	f.Add(1)
	f.Add(2)
	f.Add(3)
	fmt.Printf("%d\t%d\n", f.array, cap(f.array))
	f.Add(4)
	fmt.Printf("%d\t%d\n", f.array, cap(f.array))

	result, err := f.Remove()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%d\n", result)

	result, err = f.Remove()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%d\n", result)

	result, err = f.Remove()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%d\n", result)

	result, err = f.Remove()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%d\n", result)

	result, err = f.Remove()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%d\n", result)
}
