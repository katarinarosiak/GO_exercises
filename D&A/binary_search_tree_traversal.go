/*
Depth First Traversal:
- Inorder (Left, Root, Right)
- Preorder (Root, Left, Right)
- Postorder (Left, Right, Root)

- Breadth-First or Level Order Traveral


*/

package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func PrintInorder(root *Node) {
	/*
			 1. Traverse the left subtree, i.e., call Inorder(left-subtree)
		   2. Visit the root.
		   3. Traverse the right subtree, i.e., call Inorder(right-subtree)


	*/

	if root == nil {
		return
	}

	PrintInorder(root.Left)
	fmt.Println(root.Key)
	PrintInorder(root.Right)

}

func PrintPreorder(root *Node) {
	/*
			 1. Visit the root.
		   2. Traverse the left subtree, i.e., call Preorder(left-subtree)
		   3. Traverse the right subtree, i.e., call Preorder(right-subtree)
	*/

	if root == nil {
		return
	}

	fmt.Println(root.Key)
	PrintPreorder(root.Left)
	PrintPreorder(root.Right)
}

func PrintPostorder(root *Node) {
	/*
			 1. Traverse the left subtree, i.e., call Postorder(left-subtree)
		   2. Traverse the right subtree, i.e., call Postorder(right-subtree)
		   3. Visit the root.

	*/
	if root == nil {
		return
	}

	PrintPostorder(root.Left)
	PrintPostorder(root.Right)
	fmt.Println(root.Key)

}

func PrintlevelOrder(root *Node) [][]int {

	/*
		1. check if the root is nill
		2. create a queue and insert the root into it
		3. create a slice of slice of integers to hold final result
		4. Queue processing: as long as the queue is not empty,
		process the elements from the queue.
		(the curr size of the queue indicates total num of nodes that are
			on the same level)
			5. Process eacg node from teh same level:
			- remove the node from the queue,
			- insert its val to array that holds nodes' values from the curr level
			insert its children into the queue for later
			6. Once all the nodes from the same level are processed, the
			output of the current level is inserted into the final


	*/

	if root == nil {
		return [][]int{}
	}

	// Create a queue and insert root
	queue := []*Node{}
	queue = append(queue, root)

	// Create result slice
	result := [][]int{}

	// Process as long as queue is not empty
	for len(queue) > 0 {
		// Get the current size or length of the queue.
		// This indicates the total number of nodes that are part of the current level
		sz := len(queue)
		level := []int{}
		// Process "sz" number of elements from the queue and fill the values
		// into "level" array/slice

		for i := 0; i < sz; i++ {
			//remove node
			node := queue[0]
			queue = queue[1:]

			// Visit the node. Here visiting means collecting it into the output array
			level = append(level, node.Key)

			// Insert children of the node into the queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		//now level is filled with current level of node's vals
		//insert into final result
		result = append(result, level)
	}
	//return result
	return result
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(44)
	tree.Insert(200)
	tree.Insert(47)
	tree.Insert(33)
	tree.Insert(255)
	tree.Insert(311)

	// firstNode := tree.GetRootNote()

	PrintPostorder(tree)
	// fmt.Println("here")

	data := PrintlevelOrder(tree)
	fmt.Println(data)

}
