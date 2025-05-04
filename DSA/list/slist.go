package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}
type LinkedList struct {
	headNode *Node
}

// add to head
func (ll *LinkedList) AddToHead(property int) {
	// var node = Node{}
	// node.property = property
	// node.nextNode = ll.headNode
	// ll.headNode = &node

	var node = Node{
		property: property,
		nextNode: ll.headNode,
	}
	ll.headNode = &node

}

// iterate list
func (ll *LinkedList) IterateList() {
	for e := ll.headNode; e != nil; e = e.nextNode {
		fmt.Println(e.property)
	}

}

// last node
func (ll *LinkedList) LastNode() *Node {
	var lastNode *Node

	for node := ll.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// add to end
func (ll *LinkedList) AddToEnd(property int) {
	lastNode := ll.LastNode()
	lastNode.nextNode = &Node{
		property: property,
		nextNode: nil,
	}
}
func (ll *LinkedList) NodeWithValue(property int) *Node {
	for e := ll.headNode; e != nil; e = e.nextNode {
		if e.property == property {
			return e
		}
	}
	return &Node{}
}

// add after

func (ll *LinkedList) AddAfter(nodeProperty, property int) {

	for e := ll.headNode; e != nil; e = e.nextNode {
		if e.property == nodeProperty {
			tmp := e.nextNode
			e.nextNode = &Node{
				property: property,
				nextNode: tmp,
			}

			break
		}
	}
}
func main() {
	var ll LinkedList
	ll.AddToHead(1)
	ll.AddToHead(2)
	ll.AddToHead(3)
	ll.AddToHead(4)
	ll.AddAfter(2, 7)
	ll.IterateList()

}
