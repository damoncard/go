package tools

import (
	"fmt"
	"sync"
)

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

var (
	Root *Node
	wg   sync.WaitGroup
)

func NewNode(nums ...int) {
	for _, v := range nums {
		if Root == nil {
			Root = &Node{nil, nil, v}
		} else {
			addNode(v)
		}
	}
}

func addNode(v int) {
	var tempNode *Node = Root
	for true {
		if v < tempNode.Value {
			if tempNode.Left != nil {
				tempNode = tempNode.Left
			} else {
				tempNode.Left = &Node{nil, nil, v}
				return
			}
		} else {
			if tempNode.Right != nil {
				tempNode = tempNode.Right
			} else {
				tempNode.Right = &Node{nil, nil, v}
				return
			}
		}
	}
}

func SearchNode(num int) bool {
	var tempNode *Node = Root
	for tempNode != nil {
		if num < tempNode.Value {
			tempNode = tempNode.Left
		} else if num > tempNode.Value {
			tempNode = tempNode.Right
		} else {
			return true
		}
	}
	return false
}

func getNode(ptr *Node, c chan<- int) {
	defer wg.Done()
	c <- ptr.Value
	if ptr.Left != nil {
		go getNode(ptr.Left, c)
	} else {
		c <- -1
	}
	if ptr.Right != nil {
		go getNode(ptr.Right, c)
	} else {
		c <- -1
	}
}

func closeChannel(c chan int) {
	wg.Wait()
	close(c)
}

func PrintNode(n int) {
	c := make(chan int)
	wg.Add(n)
	go getNode(Root, c)
	go closeChannel(c)
	var nodes []int
	for i := range c {
		nodes = append(nodes, i)
	}
	fmt.Println(nodes)
}
