package structures

import (
	"errors"
	"fmt"
)

type HashTableTwo struct {
	table [arraySize]*HashNodeAdress
}

func (hashMap *HashTableTwo) Clear() error {
	for i := 0; i < arraySize; i++ {
		if hashMap.table[i] != nil {
			hashMap.table[i] = nil
		}
	}
	return nil
}

// Insert Добавление значения в Хэш-таблицу.
func (hashMap *HashTableTwo) Insert(key string, value string) error {
	keyValue := &HashNodeAdress{key: key, value: value}
	index := hashFunction(keyValue.key)

	if hashMap.table[index] == nil {
		hashMap.table[index] = keyValue
		return nil
	}

	if hashMap.table[index].key == key {
		return errors.New("such a key already exists")
	}

	for i := (index + 1) % arraySize; i != index; i = (i + 1) % arraySize {
		if i == index {
			return errors.New("table is full")
		}

		if hashMap.table[i] == nil {
			hashMap.table[i] = keyValue
			return nil
		}

		if hashMap.table[i].key == key {
			return errors.New("such a key already exists")
		}
	}

	return errors.New("failed to add element")
}

// Remove Удаление значения в Хэш-таблицы.
func (hashMap *HashTableTwo) Remove(key string) (string, error) {
	index := hashFunction(key)

	if hashMap.table[index] == nil {
		return "", errors.New("no such meaning")
	}

	// Если найден ключ в текущем индексе, удаляем его
	if hashMap.table[index].key == key {
		remove := hashMap.table[index].value
		hashMap.table[index] = nil
		return remove, nil
	}

	// В случае коллизии, ищем ключ в следующих индексах
	for i := (index + 1) % arraySize; i != index; i = (i + 1) % arraySize {
		if hashMap.table[i] != nil && hashMap.table[i].key == key {
			remove := hashMap.table[i].value
			hashMap.table[i] = nil
			return remove, nil
		}
	}

	return "", errors.New("no such meaning")
}

// Get Получение элемента хэш-таблицы по индексу.
func (hashMap *HashTableTwo) Get(key string) (string, error) {
	index := hashFunction(key)

	if hashMap.table[index] == nil {
		return "", errors.New("element not found")
	} else if hashMap.table[index].key == key {
		return hashMap.table[index].value, nil
	} else {
		for i := (index + 1) % arraySize; i != index; i = (i + 1) % arraySize {
			if hashMap.table[i] == nil {
				return "", errors.New("element not found")
			}

			if hashMap.table[i].key == key {
				return hashMap.table[i].value, nil
			}
		}
	}

	return "", errors.New("element not found")
}

// PrintToConsole Вывод содержимого Хэш-таблицы в консоль.
func (hashMap *HashTableTwo) PrintToConsole() {
	for i, node := range hashMap.table {
		if hashMap.table[i] == nil {
			continue
		}

		if node != nil {
			fmt.Printf("Index %d: (%s, %s)\n", i, node.key, node.value)
		}
	}
}
