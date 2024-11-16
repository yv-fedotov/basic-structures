package main

import (
	"basic-structures/homework"
	"fmt"
	"time"
)

func main() {

	fmt.Printf("single array:\tadd\tremove\n")
	single := homework.NewSingleArray[int](100)
	start := time.Now()
	for i := 0; i < 100; i++ {
		if err := single.Add(i, i); err != nil {
			fmt.Println(err)
			return
		}
	}
	stop := time.Since(start)
	fmt.Printf("\t\t%d\t", stop)
	start = time.Now()
	for i := 0; i < 100; i++ {
		if _, err := single.Remove(i); err != nil {
			fmt.Println(err)
		}
	}

	stop = time.Since(start)
	fmt.Printf("%d\n", stop)

	list := homework.NewLinkList[int]()
	fmt.Printf("linked list:\tpop\tpush\n")
	start = time.Now()
	for i := 0; i < 100; i++ {
		list.Push(i)
	}
	stop = time.Since(start)
	fmt.Printf("\t\t%d\t", stop)

	start = time.Now()
	for i := 0; i < 100; i++ {
		list.Pop()
	}
	stop = time.Since(start)
	fmt.Printf("%d\n", stop)

	fmt.Printf("priority queue array:\tenqueue\tdequeue\n")
	queue := homework.NewMatrixArray[int](100)
	start = time.Now()
	for i := 0; i < 100; i++ {
		queue.Enqueue(i, i)
	}
	stop = time.Since(start)
	fmt.Printf("\t\t\t%d\t", stop)

	for i := 0; i < 100; i++ {
		if _, err := queue.Dequeue(); err != nil {
			fmt.Println(err)
		}
	}
	stop = time.Since(start)
	fmt.Printf("%d\n", stop)

	fmt.Printf("priority queue linked list:\tenqueue\tdequeue\n")
	qList := homework.NewPriorityQueue[int]()
	start = time.Now()
	for i := 0; i < 100; i++ {
		qList.Enqueue(i, i)
	}
	stop = time.Since(start)
	fmt.Printf("\t\t\t\t%d\t", stop)

	for i := 0; i < 100; i++ {
		qList.Dequeue()
	}
	stop = time.Since(start)
	fmt.Printf("%d\n", stop)
}
