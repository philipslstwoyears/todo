package main

import (
	"container/list"
	"fmt"
)

func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {
	elem := queue.Front()
	queue.Remove(elem)
	return elem
}

func printQueue(queue *list.List) {
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
}

func main() {

	queue := list.New()

	Push(1, queue)

	Push(2, queue)

	Push(3, queue)

	printQueue(queue) // 123

	Pop(queue)

	printQueue(queue) // в ту же строку: 23

}
