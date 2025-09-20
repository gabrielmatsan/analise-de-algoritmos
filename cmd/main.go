package main

import (
	"fmt"

	"github.com/gabrielmatsan/analise-de-algoritmos/internal"
)

// Função principal para demonstração
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Array original:", arr)

	sortedArr := internal.MergeSort(arr)
	fmt.Println("Array ordenado:", sortedArr)
}
