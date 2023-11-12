package structures

import (
	"fmt"
)

type Position struct {
	X, Y int
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

var dx = []int{1, 2, 2, 1, -1, -2, -2, -1}
var dy = []int{2, 1, -1, -2, -2, -1, 1, 2}

func isInsideBoard(x, y, N int) bool {
	return x >= 0 && x < N && y >= 0 && y < N
}

// Enqueue Добавление значения в очередь.
func (queue *Queue) Enqueue(position Position) {
	node := &QueueNode{position: position}
	if queue.head == nil {
		queue.head = node
		queue.tail = node
	} else {
		queue.tail.next = node
		queue.tail = node
	}
}

// Dequeue Удаление значения из очереди.
func (queue *Queue) Dequeue() Position {
	if queue.head == nil {
		fmt.Println("queue is empty")
	}

	value := queue.head.position
	queue.head = queue.head.next
	if queue.head == nil {
		queue.tail = nil
	}

	return value
}

func (queue *Queue) isEmpty() bool {
	return queue.head == nil
}

func FindShortestPath(N int, knightPos, newPos Position) ([]Position, int) {
	queue := &Queue{}
	queue.Enqueue(knightPos)

	visited := make(map[Position]bool)
	parent := make(map[Position]Position)

	for !queue.isEmpty() {
		curr := queue.Dequeue()

		if curr.X == newPos.X && curr.Y == newPos.Y {
			path := []Position{curr}

			for curr != knightPos {
				curr = parent[curr]
				path = append([]Position{curr}, path...)
			}

			return path, len(path) - 1
		}

		for i := 0; i < 8; i++ {
			newX, newY := curr.X+dx[i], curr.Y+dy[i]
			nextPos := Position{newX, newY}

			if isInsideBoard(newX, newY, N) && !visited[nextPos] {
				queue.Enqueue(nextPos)
				visited[nextPos] = true
				parent[nextPos] = curr
			}

		}
	}

	return nil, -1
}
