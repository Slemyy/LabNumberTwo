package main

import (
	"LabNumberTwo/structures"
	"fmt"
)

func main() {
	// Первое задание (Стек)
	fmt.Println("\n-- (Stack) --")
	fmt.Println("({}) -", structures.IsValidHooks("({})"))
	fmt.Println("({)} -", structures.IsValidHooks("({)}"))
	fmt.Println("(][) -", structures.IsValidHooks("(][)"))
	fmt.Println("\"\" -", structures.IsValidHooks(""))
	fmt.Println("[] -", structures.IsValidHooks("[]"))

	// Второе задание (Множество)
	fmt.Println("\n-- (Set) --")
	set1 := &structures.Set{}
	err := set1.Add("a")
	if err != nil {
		return
	}

	err = set1.Add("b")
	if err != nil {
		return
	}

	set2 := &structures.Set{}
	err = set2.Add("b")
	if err != nil {
		return
	}

	err = set2.Add("c")
	if err != nil {
		return
	}

	fmt.Println(set2.Itmember("52"))

	set3 := &structures.Set{}
	err = set3.Add("c")
	if err != nil {
		return
	}

	err = set3.Add("d")
	if err != nil {
		return
	}

	result := structures.Union(set1, set2, set3)
	result.PrintToConsole()

	newDiff := structures.Difference(set1, set2)
	newDiff.PrintToConsole()

	newIter := structures.Intersection(set1, set2)
	newIter.PrintToConsole()

	// Третье задание (Массив)
	fmt.Println("\n-- (Array) --")
	array := &structures.DynamicArray{Data: []string{"x", "y", "z"}}
	fmt.Println(array.SubArray())

	// Четвертое задание (Двоичное дерево)
	fmt.Println("\n-- (Binary Search Tree) --")
	BST := &structures.BinarySearchTree{}

	BST.Insert(6)
	BST.Insert(5)
	BST.Insert(3)
	BST.Insert(16)
	BST.Insert(14)
	BST.Insert(13)
	BST.Insert(15)

	fmt.Println("Binary Search Tree:")
	BST.PrintToConsole()

	if BST.IsBST() == true {
		fmt.Println("\nThis tree is BST.")
	} else {
		fmt.Println("\nThis tree is not BST.")
	}

	// Пятое задание (Очередь)
	fmt.Println("\n-- (Queue) --")

	// Шестое задание (Хеш-таблица)
	fmt.Println("\n-- (HashTableChains) --")
	hashOne := &structures.HashTableOne{}

	err = hashOne.Insert("1", "Kirill")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashOne.Insert("1", "Egor")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashOne.Insert("44", "Sasha")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashOne.Insert("h", "Masha")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashOne.Insert("5", "Sonya")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashOne.Insert("12", "Alya")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashOne.Insert("21", "Alice")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashOne.Insert("21", "Pasha")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = hashOne.PrintToConsole()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println()

	remove, err := hashOne.Remove("h")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Удаленный элемент: %s\n", remove)

	err = hashOne.PrintToConsole()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println()

	get, err := hashOne.Get("1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Элемент: %s\n", get)

	// Шестое задание (Хеш-таблица)
	fmt.Println("\n-- (HashTableAdress) --")
	hashTwo := &structures.HashTableTwo{}

	err = hashTwo.Insert("1", "Kirill")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashTwo.Insert("1", "Egor")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashTwo.Insert("44", "Sasha")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashTwo.Insert("h", "Masha")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashTwo.Insert("5", "Sonya")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashTwo.Insert("12", "Alya")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashTwo.Insert("21", "Alice")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = hashTwo.Insert("21", "Pasha")
	if err != nil {
		fmt.Println(err.Error())
	}

	hashTwo.PrintToConsole()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println()

	remove, err = hashTwo.Remove("h")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Удаленный элемент: %s\n", remove)

	hashTwo.PrintToConsole()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println()

	get, err = hashTwo.Get("1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Элемент: %s\n", get)
}
