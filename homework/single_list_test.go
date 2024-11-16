package homework

import (
	"fmt"
	"testing"
)

func Test_SingleList(t *testing.T) {
	s := NewLinkList[int]()
	s.Push(1)
	s.Push(13)
	fmt.Printf("%v\n", s.head)
	e := s.Pop()
	fmt.Printf("%d\n", e)
	e = s.Pop()
	fmt.Printf("%d\n", e)
	fmt.Printf("%v\n", s.head)
	e = s.Pop()
	fmt.Printf("%d\n", e)
}
