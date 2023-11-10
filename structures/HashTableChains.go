package structures

import (
	"errors"
	"fmt"
)

type HashTableOne struct {
	table [arraySize]*HashNodeChains
}

func hashFunction(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}

	return hash % arraySize
}

func (hashMap *HashTableOne) Insert(key string, value string) error {
	keyValue := &HashNodeChains{key: key, value: value}
	index := hashFunction(keyValue.key)

	if hashMap.table[index] == nil {
		hashMap.table[index] = keyValue
		return nil
	}

	if hashMap.table[index].key == key {
		return errors.New("such a key already exists")
	}

	curr := hashMap.table[index]
	for curr.next != nil {
		if curr.next.key == key {
			return errors.New("such a key already exists")
		}

		curr = curr.next
	}
	curr.next = keyValue

	return nil
}

func (hashMap *HashTableOne) Remove(key string) (string, error) {
	index := hashFunction(key)

	if hashMap.table[index] == nil {
		return "", errors.New("no such meaning")
	}

	curr := hashMap.table[index]
	if curr.key == key {
		hashMap.table[index] = curr.next
		return curr.value, nil
	}

	prev := curr
	for curr != nil && curr.key != key {
		prev = curr
		curr = curr.next
	}

	if curr == nil {
		return "", nil
	}

	prev.next = curr.next
	return curr.value, nil
}

func (hashMap *HashTableOne) Get(key string) (string, error) {
	index := hashFunction(key)
	if hashMap.table[index] == nil {
		return "", errors.New("value not found")
	} else if hashMap.table[index].key == key {
		return hashMap.table[index].value, nil
	}

	curr := hashMap.table[index]

	for curr != nil {
		if curr.key == key {
			return curr.value, nil
		}
		curr = curr.next
	}

	return curr.value, nil
}

func (hashMap *HashTableOne) PrintToConsole() error {
	for i, node := range hashMap.table {
		if hashMap.table[i] == nil {
			continue
		}
		fmt.Printf("Index %d: ", i)
		curr := node
		for curr != nil {
			fmt.Printf("(%s, %s) ", curr.key, curr.value)
			curr = curr.next
		}
		fmt.Println()
	}

	return nil
}
