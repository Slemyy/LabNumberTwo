package structures

type Node struct {
	data string
	next *Node
}

type QueueNode struct {
	data int
	next *QueueNode
}

const arraySize = 769 // 769

type HashNodeChains struct {
	key   string
	value string
	next  *HashNodeChains
}

type HashNodeAdress struct {
	key   string
	value string
}

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}
