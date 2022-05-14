// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
)

func main() {
	ll := makeLinkedList(1, 4)
	printList(ll)

	reverseBetween(ll, 2, 3)

	printList(ll)
}

type ListNode struct {
	Val  int
	Next *ListNode
}


/*
temp 
left
right
dummy 

con 
tail 

prev = nil 
curr = head 
counter = 1

L-1 times iteration


for counter == right
	left => right
	right = right.Next
	counter++

	if counter == left {
		con = left
		tail = right
	}

	if counter == left+1 {

	}




*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	
	//create a dummy node
	// dummy node will help us to deal with wdgw cases for example when 
	// the list to reverse at the beginning of the list head == position 1

	var dummy *ListNode
	leftPrev, curr := dummy, head

	//reach node at position 'left'
	for i := 0; i < left -1; i++ {
		leftPrev, curr= curr, curr.Next 
	}


	//right-left + 1


	//now cur="left"
	// leftPrev="node before left"
	// reverse from left to right 

	prev := nil
	for i := 0; i < (ight - left + 1) {
		tempNext := curr.Next    //save the next to the temp
		curr.Next = prev        
		prev, curr = curr, tempNext
	}

	//updating the pointers 
	leftPrev.Next.Next = curr 
	leftPrev.next = prev 
	return dummy.Next
}





// func reverseBetween(head *ListNode, left int, right int) *ListNode {
// 	curr := head
// 	next := curr.Next
// 	var prev *ListNode
// 	// var next *ListNode
// 	var temp *ListNode
// 	var subTail *ListNode
// 	nodeCount := 1

// 	for nodeCount < left {
// 		nodeCount++
// 		prev = curr
// 		curr = next
// 		next = next.Next
// 		fmt.Println("walk")
// 	}
// 	fmt.Println("nodeCount", nodeCount)
// 	fmt.Println("curr", curr.Val, "next", next.Val)
// 	temp = prev // last node before left
// 	subTail = curr
// 	// nodeCount = walk(nodeCount, prev, curr, next)

// 	//walk manually
// 	nodeCount++
// 	prev = curr
// 	curr = next
// 	next = next.Next
// 	fmt.Println("walk")

// 	for nodeCount <= right { //start when previous is on L and stop when previous is on R
// 		fmt.Println("nodeCount", nodeCount)
// 		// fmt.Println("prev", prev.Val, "curr", curr.Val, "next", next.Val)
// 		curr.Next = prev //rewire

// 		//walk manually
// 		nodeCount++
// 		prev = curr
// 		curr = next
// 		if next != nil {
// 			next = next.Next
// 		}
// 		fmt.Println("walk")
// 	}

// 	// //point temp to last reversed
// 	temp.Next = prev
// 	subTail.Next = curr

// 	return head
// }

// func walk(nodeCount int, prev, curr, next *ListNode) int {
// 	nodeCount++
// 	prev = curr
// 	curr = next
// 	next = next.Next
// 	fmt.Println("curr", curr.Val)
// 	return nodeCount
// }

// // >> 1 4 3 2 4 5

// func printList(head *ListNode) {
// 	if head == nil {
// 		fmt.Println("Empty!")
// 	}
// 	nodeNumber := 1
// 	currNode := head
// 	for currNode != nil {
// 		fmt.Println("node #", nodeNumber, " value: ", currNode.Val)
// 		currNode = currNode.Next
// 		nodeNumber++
// 	}
// }

// func makeLinkedList(min, max int) *ListNode {
// 	currValue := min
// 	firstNode := &ListNode{
// 		Val: min,
// 	}
// 	var currNode = firstNode
// 	for currValue <= max {
// 		currValue++
// 		currNode.Next = &ListNode{
// 			Val: currValue,
// 		}
// 		currNode = currNode.Next
// 	}
// 	return firstNode
// }








