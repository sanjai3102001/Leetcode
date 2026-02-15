package main

import "fmt"

type BinaryTree struct {
	val   int
	Left  *BinaryTree
	Right *BinaryTree
}

// here I will be making some basic binary tree and preorder , inorder and postorder
func main() {
	// Creating root
	root := &BinaryTree{val: 3}
	backuproot := root

	root.Left = &BinaryTree{val: 9}
	root.Right = &BinaryTree{val: 20}
	root = root.Right
	root.Left = &BinaryTree{val: 15}
	root.Right = &BinaryTree{val: 7}

	root = backuproot
	fmt.Println("Inorder")
	Inorder(root)
	fmt.Println("Postorder")
	Postorder(root)
	fmt.Println("Preorder")
	preorder(root)

}

func Inorder(root *BinaryTree) {

	if root == nil {
		return
	}
	Inorder(root.Left)
	fmt.Println(root.val)
	Inorder(root.Right)

}

func Postorder(root *BinaryTree) {
	if root == nil {
		return
	}
	Postorder(root.Left)
	Postorder(root.Right)
	fmt.Println(root.val)
}

func preorder(root *BinaryTree) {
	if root == nil {
		return
	}
	fmt.Println(root.val)
	preorder(root.Left)
	preorder(root.Right)
}
