package structures

import (
	"fmt"
	"math"
)

type DynamicArray struct {
	Data   []string
	length int
}

func (array *DynamicArray) AddToSubset(subset []string, item string) []string {
	newSubset := make([]string, len(subset)+1)
	newSubset[0] = item
	copy(newSubset[1:], subset)
	return newSubset
}

func (array *DynamicArray) AddToResult(result [][]string, subset []string) [][]string {
	newResult := make([][]string, len(result)+1)

	for k := 0; k < len(result); k++ {
		newResult[k] = result[k]
	}

	newResult[len(result)] = subset
	return newResult
}

func (array *DynamicArray) GetItem(index int) (bool, string) {
	if index >= 0 && index < len(array.Data) {
		return true, array.Data[index]
	}
	return false, ""
}

func (array *DynamicArray) SubArray() [][]string {
	n := len(array.Data)
	result := make([][]string, 0, int(math.Pow(2, float64(n))))

	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
		subset := make([]string, 0, n)
		for j := 0; j < n; j++ {
			if i&(int(math.Pow(2, float64(j)))) > 0 {
				fmt.Printf("(%d, %d) ", i, j)
				_, item := array.GetItem(j)
				subset = array.AddToSubset(subset, item)
			}
		}

		result = array.AddToResult(result, subset)
	}

	return result
}
