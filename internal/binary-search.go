package internal

import "fmt"

func BinarySearch(arr []int, target int) (int, error) {

	// Contagem: 2 operações de atribuição, 1 operação de subtração e 1 operação de tamanho do array = 4,
	left, right := 0, len(arr)-1 // C₁ = 4

	// Contagem: n operações, log₂(n)
	for left <= right { // C₂ x log₂(n)
		// Contagem: 3 operações de atribuição (soma, divisão e atribuição)
		mid := (left + right) / 2 // C₃ x log₂(n)

		// 1 acesso de array + 1 comparação de target = 2, C₄ x log₂(n)
		if arr[mid] == target { // C₄ x log₂(n)
			return mid, nil // C₅ x 1 (apenas um retorno)
		}

		// se o elemento do meio for menor que o target, o target está na metade direita
		// 1 acesso de array + 1 comparação de target = 2, C₆ x log₂(n)
		if arr[mid] < target {
			// Contagem: 1 operação de atribuição

			// 1 soma e 1 atribuição = 2, C₇ x log₂(n/2)
			left = mid + 1
		} else {
			// Contagem: 1 soma e 1 atribuição = 2, C₈ x log₂(n/2)
			right = mid - 1
		}
	}
	// Contagem: 1 operação de retorno
	return -1, fmt.Errorf("target not found") // C₉ x 1
}

//f(n) = C₁ + (C₂ + C₃ + C₄ + C₆) × log₂(n) + (C₇ + C₈) × log₂(n)/2 + C₉
//Simplificando: f(n) = c₁ + c₂ × log₂(n)

// Melhor caso: encontra o target na primeira iteração, f(n) = O(n)
// Pior caso: não encontra o target, f(n) = O(log(n))
// Caso médio: f(n) = O(log(n))

// Big-O: O(log(n))
// Big-Omega: O(1)
// Big-Theta: O(log(n))
