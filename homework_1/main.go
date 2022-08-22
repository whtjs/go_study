package main

import (
	"fmt"
	"homework_1/DS"
)

func main() {

	StackT()
	QueueT()
}

func StackT() {
	s := DS.NewStack[int]()
	s.Push(1)
	s.Push(4)

	fmt.Println(s.Get())
	fmt.Println(s.Pop())
	fmt.Println(s.Get())
}

func QueueT() {
	q := DS.Queue[string]{}

	q.Enqueue("hi")
	q.Enqueue("hasun")

	fmt.Println(q.Get())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Get())

}