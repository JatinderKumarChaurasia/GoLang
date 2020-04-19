package main

import (
	"fmt"
	"math/rand"
	"time"
)

type linkedList struct {
	Value int
	Node  *linkedList
}

func CreateLinkedList(num int, temp *linkedList) *linkedList {
	if temp == nil {
		temp = &linkedList{
			Value: rand.Intn(100),
			Node:  nil,
		}
		num--
	}
	tempList := temp
	for i := 0; i < num; i++ {
		tem := &linkedList{
			Value: rand.Intn(100),
			Node:  nil,
		}
		tempList.Node = tem
		tempList = tempList.Node
	}
	return temp
}

func main() {
	// LinkedList.go
	fmt.Println("Generating Linked List ......................")
	rand.Seed(time.Now().UnixNano())
	list := CreateLinkedList(10, nil)

	for ; list != nil; list = list.Node {
		fmt.Println("Printing Request", list.Value)

	}
	fmt.Println("Generating Linked List Completed......................")
}
