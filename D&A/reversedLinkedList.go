package main

import (
	"fmt"
)

type Node struct {
	data     int
	previous *Node
	next     *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *LinkedList) Prepend(n *Node) {
	if l.length == 0 {
		l.head = n
		l.tail = n
		l.head.next = nil
		l.head.previous = nil
	} else {
		second := l.head
		l.head = n
		l.head.next = second
		l.head.next.previous = n
	}
	l.length++

}

func (l *LinkedList) Append(n *Node) {
	if l.length == 0 {
		l.head = n
		l.tail = n
	} else {
		temp := l.tail
		temp.next = n
		l.tail = n
		l.tail.previous = temp

	}
	l.length++
}

func (l LinkedList) PrintListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *LinkedList) DeleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := LinkedList{}
	node1 := &Node{data: 1}
	node2 := &Node{data: 2}
	node3 := &Node{data: 3}
	myList.Prepend(node1)
	myList.Prepend(node2)
	myList.Prepend(node3)

	node4 := &Node{data: 4}
	node5 := &Node{data: 41}
	myList.Append(node4)
	myList.Append(node5)

	myList.PrintListData()
}
