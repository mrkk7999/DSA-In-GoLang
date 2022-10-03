package main

import "fmt"

// Structure
// Node -> individual node of linked list
type node struct {
	nodeVal  interface{}
	nextNode *node
}

func initEmptyList() *node {
	return &node{
		nodeVal:  nil,
		nextNode: nil,
	}
}

func initWithHead(ele interface{}) *node {
	return &node{
		nodeVal:  ele,
		nextNode: nil,
	}
}

func main() {
	singleList := initEmptyList()
	singleList = initWithHead("Name :- ")
	fmt.Println(singleList.nodeVal)
	//tempNode := singleList
	//// tempNode.nodeVal = "Mr."
	//tempNode.nextNode = initWithHead("Mr.")
	//tempNode = tempNode.nextNode
	//tempNode.nextNode = initWithHead("Kiran")
	//tempNode = tempNode.nextNode
	//tempNode.nextNode = initWithHead("Arun")
	//tempNode = tempNode.nextNode
	//tempNode.nextNode = initWithHead("Kshirsagar")
	//tempNode = tempNode.nextNode
	//temp := singleList
	//for temp.nextNode != nil {
	//	fmt.Print(temp.nodeVal, "\t")
	//	temp = temp.nextNode
	//}
	//fmt.Print(temp.nodeVal)
}
