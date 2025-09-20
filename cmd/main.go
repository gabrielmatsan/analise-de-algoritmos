package main

import (
	"fmt"

	"github.com/gabrielmatsan/analise-de-algoritmos/internal"
)

// Função principal para demonstração
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Array original:", arr)

	//sortedArr := internal.MergeSort(arr)
	sortedArr := internal.QuickSort(arr)
	fmt.Println("Array ordenado:", sortedArr)

	target := 25
	index, err := internal.BinarySearch(sortedArr, target)
	if err != nil {
		fmt.Println("Target not found")

	}
	fmt.Println("Target found at index:", index)

}
