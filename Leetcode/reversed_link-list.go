/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }

https://leetcode.com/problems/reverse-linked-list/submissions/
*/

package main

func main() {
	// reverseList(list)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var temp *ListNode
	var left *ListNode
	right := head

	for right != nil {
		temp = right.Next
		right.Next = left

		left = right
		right = temp
	}

	return left
}
