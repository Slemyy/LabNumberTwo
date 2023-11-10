package structures

import (
	"errors"
)

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

// Enqueue Добавление значения в очередь.
func (queue *Queue) Enqueue(value int) {
	node := &QueueNode{data: value}
	if queue.head == nil {
		queue.head = node
		queue.tail = node
	} else {
		queue.tail.next = node
		queue.tail = node
	}
}

// Dequeue Удаление значения из очереди.
func (queue *Queue) Dequeue() (int, error) {
	if queue.head == nil {
		return 0, errors.New("queue is empty")
	} else {
		value := queue.head.data
		queue.head = queue.head.next
		if queue.head == nil {
			queue.tail = nil
		}

		return value, nil
	}
}
