package structures

import (
	"fmt"
	"math"
)

type BinarySearchTree struct {
	root *TreeNode
}

func (bst *BinarySearchTree) Insert(value int) {
	newNode := &TreeNode{Data: value}

	if bst.root == nil {
		bst.root = newNode
		return
	} else {
		insertNode(bst.root, newNode)
	}
}

func insertNode(node, newNode *TreeNode) {
	if newNode.Data < node.Data {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNode(node.Left, newNode)
		}
	} else if newNode.Data > node.Data {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode(node.Right, newNode)
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

	if value < root.Data {
		root.Left = deleteNode(root.Left, value)
	} else if value > root.Data {
		root.Right = deleteNode(root.Right, value)
	} else {
		// Узел с одним или нулем детей
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Узел с двумя детьми, находим наименьший узел в правом поддереве
		minRight := findMinNode(root.Right)
		root.Data = minRight.Data
		root.Right = deleteNode(root.Right, minRight.Data)
	}

	return root
}

func findMinNode(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
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
	if value < node.Data {
		return search(node.Left, value)
	} else if value > node.Data {
		return search(node.Right, value)
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
	if node.Data < min || node.Data > max {
		return false
	}

	// Рекурсивно проверяем левое и правое поддерево
	return isBST(node.Left, min, node.Data-1) && isBST(node.Right, node.Data+1, max)
}

func TestBST(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	// Проверяем, что текущий узел находится в правильном диапазоне
	if node.Data < min || node.Data > max {
		return false
	}

	// Рекурсивно проверяем левое и правое поддерево
	return isBST(node.Left, min, node.Data-1) && isBST(node.Right, node.Data+1, max)
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
		fmt.Println(node.Data)

		printTree(node.Left, indent, node.Right != nil, false)
		printTree(node.Right, indent, hasRightSibling, true)
	}
}

func Tree(node *TreeNode, indent string, hasRightSibling bool, last bool) {
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
		fmt.Println(node.Data)

		printTree(node.Left, indent, node.Right != nil, false)
		printTree(node.Right, indent, hasRightSibling, true)
	}
}
