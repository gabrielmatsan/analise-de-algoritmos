package internal

func QuickSort(arr []int) []int {
	// Esta verificação acontece UMA vez por chamada recursiva
	// Número total de chamadas: depende da árvore de recursão
	if len(arr) <= 1 { // C₁ × 1 (por chamada)
		return arr // C₂ × 1 (casos base)
	}

	// Operações de inicialização: UMA vez por chamada recursiva
	pivot := arr[len(arr)-1] // C₃ × 1 (por chamada)
	left := []int{}          // C₄ × 1 (por chamada)
	right := []int{}         // C₅ × 1 (por chamada)

	// Loop de particionamento: n-1 iterações por chamada
	// onde n = tamanho do array atual
	for i := 0; i < len(arr)-1; i++ { // C₆ × (n-1)
		if arr[i] < pivot { // C₇ × (n-1)
			left = append(left, arr[i]) // C₈ × k₁
		} else {
			right = append(right, arr[i]) // C₉ × k₂
		}
	} // k₁ + k₂ = n-1

	// Chamadas recursivas + concatenação
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
	// Custo: T(k₁) + T(k₂) + C₁₀ × n (concatenação)
}

// Melhor caso: O(n log n), quando o array já está balanceado
// Caso médio: O(n log n)
// Pior caso: O(n²), quando o array está ordenado de forma decrescente
