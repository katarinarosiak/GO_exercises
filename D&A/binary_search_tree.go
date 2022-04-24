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

func (n *Node) Search(k int) bool {
	/*
		get a curret node
		check the val of current
		if val == k return ??
		else
			if val < k
				Search(n.Left)
			else if val > k
				Search(n.Right)
	*/

	if n == nil {
		return false
	} else {
		if n.Key < k {
			return n.Right.Search(k)
		} else if n.Key > k {
			return n.Left.Search(k)
		}
	}
	return true
}

func (n *Node) TraverseAndPrint(node *Node) {
	if n == nil {
		return
	}
	if node.Left != nil {
		node.TraverseAndPrint(node.Left)
	}
	fmt.Println(n.Key)
	if node.Right != nil {
		node.TraverseAndPrint(node.Right)
	}

}

func (n *Node) findGreatest(greatest int) int {
	/*
		if n == nil
			return greatest

		if n.Right.Key > greatest
			call n.Right.findGreatest(n.Right.Key)

		return greatest
	*/
	if n == nil || n.Right == nil {
		return greatest
	}
	return n.Right.findGreatest(n.Right.Key)

}

func (n *Node) delete(k int) {
	/*
		if n has no children delete it

		if n has one child
			 delete the node
			 plug the child into the spot of deleted node

		if n has two children
			replace deleted node with successor node
				- find successor node
					visit the right child of delted value
					then keep on visiting the left until
					n has no more children
					chekc if the successor node has right child
						if so take the former right child of the successor node
							turn it into the left child of the former parent
							of the successor node





	*/
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

	fmt.Println(tree.findGreatest(tree.Key))

}
