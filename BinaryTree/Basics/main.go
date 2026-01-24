package main

import "fmt"

type BinaryTree struct {
	val   int
	Left  *BinaryTree
	Right *BinaryTree
}

// here I will be making some basic binary trii and indorder traversal
func main() {
	// Creating root
	root := &BinaryTree{val: 3}
	backuproot := root

	root.Left = &BinaryTree{val: 9}
	root.Right = &BinaryTree{val: 20}
	root = root.Right
	root.Left = &BinaryTree{val: 15}
	root.Right = &BinaryTree{val: 7}

	// for i := 1; i <= 10; i++ {
	// 	root.Left = &BinaryTree{val: i}
	// 	root = root.Left
	// }
	// root = backuproot
	// for j := 11; j <= 20; j++ {
	// 	root.Right = &BinaryTree{val: j}
	// 	root = root.Right
	// }
	root = backuproot
	fmt.Println(root.val)
	Inorder(root)

}

func Inorder(root *BinaryTree) {

	if root == nil {
		return
	}
	Inorder(root.Left)
	fmt.Println(root.val)
	Inorder(root.Right)

}
