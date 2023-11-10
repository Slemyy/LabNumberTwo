package structures

import (
	"fmt"
	"math"
)

type BinarySearchTree struct {
	root *TreeNode
}

func (bst *BinarySearchTree) Insert(value int) {
	newNode := &TreeNode{data: value}

	if bst.root == nil {
		bst.root = newNode
		return
	} else {
		insertNode(bst.root, newNode)
	}
}

func insertNode(node, newNode *TreeNode) {
	if newNode.data < node.data {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else if newNode.data > node.data {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (bst *BinarySearchTree) Delete(value int) {
	bst.root = deleteNode(bst.root, value)
}

func deleteNode(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return root
	}

	if value < root.data {
		root.left = deleteNode(root.left, value)
	} else if value > root.data {
		root.right = deleteNode(root.right, value)
	} else {
		// Узел с одним или нулем детей
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		// Узел с двумя детьми, находим наименьший узел в правом поддереве
		minRight := findMinNode(root.right)
		root.data = minRight.data
		root.right = deleteNode(root.right, minRight.data)
	}

	return root
}

func findMinNode(node *TreeNode) *TreeNode {
	for node.left != nil {
		node = node.left
	}

	return node
}

func (bst *BinarySearchTree) Search(value int) bool {
	return search(bst.root, value)
}

func search(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	if value < node.data {
		return search(node.left, value)
	} else if value > node.data {
		return search(node.right, value)
	}
	return true
}

func (bst *BinarySearchTree) IsBST() bool {
	return isBST(bst.root, math.MinInt, math.MaxInt)
}

func isBST(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	// Проверяем, что текущий узел находится в правильном диапазоне
	if node.data < min || node.data > max {
		return false
	}

	// Рекурсивно проверяем левое и правое поддерево
	return isBST(node.left, min, node.data-1) && isBST(node.right, node.data+1, max)
}

func (bst *BinarySearchTree) PrintToConsole() {
	printTree(bst.root, "", true, true)
}

func printTree(node *TreeNode, indent string, hasRightSibling bool, last bool) {
	if node != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("┗╺")
			indent += "  "
		} else {
			if hasRightSibling {
				fmt.Print("┣╸")
				indent += "┃ "
			} else {
				fmt.Print("┗╺")
				indent += "   "
			}
		}
		fmt.Println(node.data)

		printTree(node.left, indent, node.right != nil, false)
		printTree(node.right, indent, hasRightSibling, true)
	}
}
