package main

import "fmt"

// Two structures
// 1. Node -> individual node of linked list
type node struct {
	nodeVal  interface{}
	nextNode *node
}

// 2. Linked list structure -> contains list of linked nodes
type linkedList struct {
	lenOfLL int
	head    *node
}

func initEmptyList() *linkedList {
	return &linkedList{
		lenOfLL: 0,
		head: &node{
			nodeVal:  nil,
			nextNode: nil,
		},
	}
}

func initWithHead(ele interface{}) *linkedList {
	return &linkedList{
		lenOfLL: 1,
		head: &node{
			nodeVal:  ele,
			nextNode: nil,
		},
	}
}

func main() {
	singleList := initEmptyList()
	list := singleList.head
	for i := 0; i < 10; i++ {
		list.nodeVal = i
		list = list.nextNode
	}
	list = singleList.head
	for i := 0; i < 10; i++ {
		fmt.Println(list.nodeVal)
	}
}
