package internal

import "fmt"

func BinarySearch(arr []int, target int) (int, error) {

	left, right := 0, len(arr)-1 // C₁ = 4 operações

	for left <= right { // log(n), pois o espaço é dividido pela metade a cada iteração

		mid := (left + right) / 2 // 3 * log(n) operações

		if arr[mid] == target { // 2 * log(n) operações
			return mid, nil // 1 operação
		}

		if arr[mid] < target { // 2 * log(n) operações
			left = mid + 1 // 1 * log(n) operações
		} else {
			right = mid - 1 // 1 * log(n) operações
		}
	}
	return -1, fmt.Errorf("target not found") // 1 operação
}

// Função de Complexidade:
// T(n) = 5 + 10 * log(n)

// Big-O: O(log(n)) - Pior caso
// Big-Omega: O(1) - Melhor caso
// Big-Theta: O(log(n)) - Caso médio
