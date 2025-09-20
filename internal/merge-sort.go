package internal

func MergeSort(arr []int) []int {

	// Contagem: 1 operação (comparar)
	if len(arr) <= 1 {
		return arr
	}

	// Contagem: 1 operação (dividir)
	mid := len(arr) / 2

	// Contagem: T(n/2) + T(n/2) operações para cada chamada recursiva, alem de atribuir os valores de left e right, 1 operação cada
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Contagem: n operações para mesclar os dois arrays ordenados
	return merge(left, right)
}

func merge(left, right []int) []int {
	//merge: intercal dois arrays ordenados em um único array ordenado, min(len(left), len(right)) operações
	result := make([]int, 0, len(left)+len(right))

	// Contagem: 1 operação para cada
	i, j := 0, 0

	// Contage: 3nk
	for i < len(left) && len(right) > j {
		// Contagem: nk operações
		if left[i] <= right[j] {
			// Contagem: 1,2,3
			result = append(result, left[i])
			// Contagem: 1
			i++
		} else {
			// Contagem: 1,2,3
			result = append(result, right[j])
			// Contagem: 1
			j++
		}
	}

	// Contagem: nk operações
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Contagem: nk operações
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	// 1 operação
	return result
}
