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

// Função de Complexidade:
// T(n) = {
//   C₁ + C₂,                           se n ≤ 1 (caso base)
//   C₃ + C₄ + C₅ + (C₆ + C₇)(n-1) +   se n > 1
//   C₈×k₁ + C₉×k₂ + T(k₁) + T(k₂) + C₁₀×n
// }
//
// Onde: k₁ + k₂ = n-1 (elementos particionados)
//
// ANÁLISE POR CASOS:
//
// 1. MELHOR CASO: Partição sempre balanceada (k₁ ≈ k₂ ≈ n/2)
//    T(n) = 2T(n/2) + Θ(n)
//    Pelo Teorema Mestre: T(n) = Θ(n log n)
//
// 2. CASO MÉDIO: Partições em média balanceadas
//    T(n) = Θ(n log n)
//
// 3. PIOR CASO: Partição sempre desbalanceada (k₁ = n-1, k₂ = 0 ou vice-versa)
//    T(n) = T(n-1) + T(0) + Θ(n) = T(n-1) + Θ(n)
//    Resolvendo: T(n) = Θ(n²)
//
// COMPLEXIDADE DE ESPAÇO:
// - Melhor/Médio caso: Θ(log n) - profundidade da recursão
// - Pior caso: Θ(n) - quando a recursão é linear
