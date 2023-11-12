package structures

import (
	"errors"
	"fmt"
)

type Set struct {
	table [arraySize]string
}

func (set *Set) Clear() error {
	for i := 0; i < arraySize; i++ {
		if set.table[i] != "" {
			set.table[i] = ""
		}
	}
	return nil
}

// Add Добавление значения во множество.
func (set *Set) Add(value string) error {
	index := hashFunction(value)

	if set.table[index] == "" {
		set.table[index] = value
		return nil
	}

	if set.table[index] == value {
		return errors.New("such a key already exists")
	}

	for i := (index + 1) % arraySize; i != index; i = (i + 1) % arraySize {
		if i == index {
			return errors.New("table is full")
		}

		if set.table[i] == "" {
			set.table[i] = value
			return nil
		}

		if set.table[i] == value {
			return errors.New("such a key already exists")
		}
	}

	return errors.New("failed to add element")
}

// Remove Rem Удаление значения во множестве.
func (set *Set) Remove(value string) (string, error) {
	index := hashFunction(value)

	if set.table[index] == "" {
		return "", errors.New("no such meaning")
	}

	if set.table[index] == value {
		remove := set.table[index]
		set.table[index] = ""
		return remove, nil
	}

	for i := (index + 1) % arraySize; i != index; i = (i + 1) % arraySize {
		if set.table[i] == value && set.table[i] != "" {
			remove := set.table[i]
			set.table[i] = ""
			return remove, nil
		}
	}

	return "", errors.New("failed to delete element")
}

func (set *Set) Itmember(value string) (bool, error) {
	index := hashFunction(value)

	if set.table[index] == "" {
		return false, errors.New("no such meaning")
	}

	if set.table[index] == value {
		return true, nil
	}

	for i := (index + 1) % arraySize; i != index; i = (i + 1) % arraySize {
		if set.table[index] == "" {
			return false, errors.New("no such meaning")
		}

		if set.table[i] == value && set.table[i] != "" {
			return true, nil
		}
	}

	return false, errors.New("element not found")
}

// Intersection  Пересечение множеств
func Intersection(sets ...*Set) *Set {
	if len(sets) == 0 {
		return nil
	}

	interSet := Set{}
	firstSet := sets[0]

	for _, value := range firstSet.table {
		if value != "" && allSetsContainValue(value, sets[1:]) {
			err := interSet.Add(value)
			if err != nil {
				return nil
			}
		}
	}

	return &interSet
}

func allSetsContainValue(value string, sets []*Set) bool {
	for _, set := range sets {
		find, _ := set.Itmember(value)
		if !find {
			return false
		}
	}
	return true
}

// Difference Разность множеств
func Difference(sets ...*Set) *Set {
	if len(sets) == 0 {
		return nil
	}

	diffSet := Set{}
	firstSet := sets[0]

	for _, value := range firstSet.table {
		if value != "" && !anySetContainsValue(value, sets[1:]) {
			err := diffSet.Add(value)
			if err != nil {
				return nil
			}
		}
	}

	return &diffSet
}

func anySetContainsValue(value string, sets []*Set) bool {
	for _, set := range sets {
		find, _ := set.Itmember(value)
		if find {
			return true
		}
	}
	return false
}

// Union возвращает новое множество, которое является объединением других множеств.
func Union(sets ...*Set) Set {
	unionSet := Set{}

	for _, set := range sets {
		for _, value := range set.table {
			if value != "" {
				err := unionSet.Add(value)
				if err != nil {
					_ = fmt.Errorf("error: %d", err)
				}
			}
		}
	}

	return unionSet
}

// PrintToConsole Выводит множество.
func (set *Set) PrintToConsole() {
	fmt.Print("{ ")

	for i := 0; i < len(set.table); i++ {
		value := set.table[i]
		if value != "" {
			fmt.Print(value)
		}

		if i < arraySize-1 && value != "" {
			fmt.Print(" ")
		}
	}

	fmt.Println("}")
}
