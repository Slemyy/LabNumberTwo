package structures

import (
	"errors"
	"strings"
)

type Stack struct {
	head *Node
}

func (stack *Stack) Push(value string) {
	node := &Node{data: value}

	if stack.head == nil {
		stack.head = node
	} else {
		node.next = stack.head
		stack.head = node
	}
}

func (stack *Stack) Pop() (string, error) {
	if stack.head == nil {
		return "", errors.New("stack is empty")
	}

	value := stack.head.data
	stack.head = stack.head.next

	return value, nil
}

func IsValidHooks(line string) bool {
	stack := Stack{}

	openingBrackets := "({["
	closingBrackets := ")}]"

	for _, char := range line {
		charStr := string(char)
		if strings.Contains(openingBrackets, charStr) {
			stack.Push(charStr)
		} else if strings.Contains(closingBrackets, charStr) {
			value, err := stack.Pop()
			if err != nil {
				return false // Не хватает открывающей скобки
			}
			if (charStr == ")" && value != "(") ||
				(charStr == "}" && value != "{") ||
				(charStr == "]" && value != "[") {
				return false // Скобки не совпадают
			}
		}
	}

	return stack.head == nil // Если стек пуст, то скобки правильно вложены
}
