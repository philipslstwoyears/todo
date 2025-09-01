package main

import (
	"container/list"
	"fmt"
)

// ReverseList - функция для реверса списка
func ReverseList(l *list.List) *list.List {
	reversedList := list.New()
	for e := l.Front(); e != nil; e = e.Next() {
		reversedList.PushFront(e.Value)
	}
	return reversedList
}

func main() {

	checkList := list.New()

	for i := 0; i < 10; i++ {

		checkList.PushBack(i)

	}

	reversedList := (ReverseList(checkList))

	for e := reversedList.Front(); e != nil; e = e.Next() {

		fmt.Print(e.Value, " ")

	}

}
