package internal

// MergeSort implementa o algoritmo de ordenação Merge Sort
// Complexidade Temporal: O(n log n)
// Complexidade Espacial: O(n)
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

	// Contagem: O(n) operações para mesclar os dois arrays ordenados
	return merge(left, right)
}
func merge(left, right []int) []int {
	// merge: intercala dois arrays ordenados em um único array ordenado
	// Complexidade: O(n) onde n = len(left) + len(right)
	result := make([]int, 0, len(left)+len(right))

	// Contagem: 1 operação para cada => 2
	i, j := 0, 0

	// Contagem: O(n) - loop principal que percorre todos os elementos
	for i < len(left) && len(right) > j {
		// Contagem: O(1) por iteração - comparação e append
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

	// Contagem: O(n) - adiciona elementos restantes do array left
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Contagem: O(n) - adiciona elementos restantes do array right
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	// 1 operação
	return result
}

/*
ANÁLISE DE COMPLEXIDADE DO MERGE SORT:

O Merge Sort é um algoritmo de ordenação baseado na estratégia "dividir e conquistar".
Sua principal característica é ter complexidade temporal O(n log n) em todos os casos,
tornando-o muito eficiente e previsível.

CARACTERÍSTICAS PRINCIPAIS:
- Complexidade Temporal: O(n log n) - sempre, independente do input
- Complexidade Espacial: O(n) - requer espaço adicional para arrays temporários
- Algoritmo Estável: mantém a ordem relativa de elementos iguais
- Não é in-place: precisa de espaço extra para funcionar

VANTAGENS:
✓ Performance consistente O(n log n) em todos os casos
✓ Algoritmo estável
✓ Adequado para grandes volumes de dados
✓ Fácil de implementar e entender

DESVANTAGENS:
✗ Requer espaço adicional O(n)
✗ Não é in-place
✗ Pode ser mais lento que Quick Sort em casos práticos devido às constantes

O Merge Sort é especialmente útil quando:
- A estabilidade é importante
- O espaço não é uma limitação crítica
- Precisamos de performance garantida O(n log n)
- Trabalhamos com listas ligadas (onde pode ser implementado in-place)

ANÁLISE DETALHADA DA COMPLEXIDADE:

1. COMPLEXIDADE TEMPORAL O(n log n):

   A) DIVISÃO (Divide):
   - Dividir o array em duas metades: O(1)
   - Operação constante independente do tamanho

   B) CONQUISTA (Conquer):
   - Duas chamadas recursivas: T(n/2) + T(n/2) = 2T(n/2)
   - Cada subproblema tem metade do tamanho original

   C) COMBINAÇÃO (Combine):
   - Função merge: O(n)
   - Percorre todos os n elementos exatamente uma vez

   EQUAÇÃO DE RECORRÊNCIA:
   T(n) = 2T(n/2) + O(n)

   RESOLUÇÃO PELO TEOREMA MESTRE:
   - a = 2, b = 2, f(n) = n
   - n^log_b(a) = n^log_2(2) = n^1 = n
   - Como f(n) = n = n^log_b(a), temos o Caso 2
   - T(n) = O(n log n)

2. COMPLEXIDADE ESPACIAL O(n):

   A) PILHA DE RECURSÃO:
   - Altura da árvore: log₂(n)
   - Cada nível usa O(1) espaço adicional
   - Total da pilha: O(log n)

   B) ARRAYS TEMPORÁRIOS:
   - Função merge cria array result de tamanho n
   - Máximo de n elementos em arrays temporários
   - Total: O(n)

   COMPLEXIDADE ESPACIAL TOTAL: O(n)

3. ANÁLISE DO MELHOR, MÉDIO E PIOR CASO:

   - MELHOR CASO: O(n log n) - mesmo com array já ordenado
   - CASO MÉDIO: O(n log n) - comportamento consistente
   - PIOR CASO: O(n log n) - sempre divide e combina

4. COMPARAÇÃO COM OUTROS ALGORITMOS:

   Merge Sort vs Quick Sort:
   - Merge: O(n log n) garantido, mas O(n) espaço
   - Quick: O(n log n) médio, O(n²) pior caso, O(log n) espaço

   Merge Sort vs Heap Sort:
   - Merge: O(n log n) sempre, O(n) espaço, estável
   - Heap: O(n log n) sempre, O(1) espaço, não estável

5. CONSTANTES OCULTAS:

   Embora ambos sejam O(n log n), Merge Sort tem constantes maiores:
   - Mais operações de cópia de elementos
   - Overhead de chamadas recursivas
   - Alocação de memória para arrays temporários

   Por isso, em implementações práticas, Quick Sort pode ser mais rápido
   para arrays pequenos, mas Merge Sort é mais previsível e estável.
*/
