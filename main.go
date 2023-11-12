package main

import (
	"LabNumberTwo/structures"
	"fmt"
	"math"
)

func main() {
	// Первое задание (Стек)
	fmt.Println("\n-- (Stack) --")
	fmt.Println("({}) -", structures.IsValidHooks("({})"))
	fmt.Println("({)} -", structures.IsValidHooks("({)}"))
	fmt.Println("(][) -", structures.IsValidHooks("(][)"))
	fmt.Println("\"\" -", structures.IsValidHooks(""))
	fmt.Println("[] -", structures.IsValidHooks("[]"))
	fmt.Println("(()))) -", structures.IsValidHooks("(())))"))
	fmt.Println("((()) -", structures.IsValidHooks("((())"))

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
	fmt.Println("\n", array.SubArray())

	// Четвертое задание (Двоичное дерево)
	fmt.Println("\n-- (Binary Search Tree) --")
	BST := &structures.BinarySearchTree{}

	//test := &structures.TreeNode{Data: 6, Left: &structures.TreeNode{Data: 10}, Right: &structures.TreeNode{Data: 5}}
	//structures.Tree(test, "", true, true)

	test := &structures.TreeNode{Data: 8, Left: &structures.TreeNode{Data: 3, Right: &structures.TreeNode{Data: 6}}}
	structures.Tree(test, "", true, true)

	fmt.Println(structures.TestBST(test, math.MinInt, math.MaxInt))
	fmt.Println()

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
	N := 8 // Размер доски NxN
	knightPos := structures.Position{X: 1, Y: 4}
	newPos := structures.Position{X: 3, Y: 1}

	path, pathLength := structures.FindShortestPath(N, knightPos, newPos)
	if path != nil {
		fmt.Printf("Кратчайший путь: %d хода(-ов)\n", pathLength)
		fmt.Print("Координаты пути: ")

		for i, pos := range path {
			if len(path)-1 == i {
				fmt.Printf("(%d, %d).", pos.X, pos.Y)
				continue
			}

			fmt.Printf("(%d, %d) -> ", pos.X, pos.Y)
			i++
		}
	} else {
		fmt.Println("Путь не найден.")
	}

	fmt.Println()

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
