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

/*
ANÁLISE DE COMPLEXIDADE DO QUICK SORT:

O Quick Sort é um algoritmo de ordenação baseado na estratégia "dividir e conquistar"
que utiliza particionamento para dividir o array. Sua complexidade varia significativamente
dependendo da escolha do pivô e da distribuição dos dados.

CARACTERÍSTICAS PRINCIPAIS:
- Complexidade Temporal: O(n log n) médio, O(n²) pior caso
- Complexidade Espacial: O(log n) médio, O(n) pior caso
- Algoritmo In-place: pode ser implementado sem espaço extra
- Não é Estável: pode alterar a ordem relativa de elementos iguais

VANTAGENS:
✓ Performance excelente O(n log n) na prática
✓ Algoritmo in-place (versão otimizada)
✓ Constantes menores que Merge Sort
✓ Cache-friendly em implementações otimizadas

DESVANTAGENS:
✗ Pior caso O(n²) - pode ser muito lento
✗ Não é estável
✗ Performance depende da escolha do pivô
✗ Não é determinístico

ANÁLISE DETALHADA DA COMPLEXIDADE:

1. COMPLEXIDADE TEMPORAL:

   A) MELHOR CASO O(n log n):
   - Pivô sempre divide o array em duas partes iguais
   - Equação: T(n) = 2T(n/2) + O(n)
   - Resolução: T(n) = O(n log n)
   - Exemplo: array já ordenado com pivô mediano

   B) CASO MÉDIO O(n log n):
   - Pivô divide o array em proporções razoáveis
   - Análise probabilística mostra O(n log n)
   - Na prática, muito próximo do melhor caso

   C) PIOR CASO O(n²):
   - Pivô sempre é o menor ou maior elemento
   - Equação: T(n) = T(n-1) + O(n)
   - Resolução: T(n) = O(n²)
   - Exemplo: array ordenado decrescentemente

2. COMPLEXIDADE ESPACIAL:

   A) PILHA DE RECURSÃO:
   - Melhor/Médio caso: O(log n) - árvore balanceada
   - Pior caso: O(n) - árvore linear

   B) VERSÃO IN-PLACE:
   - Usa apenas O(log n) espaço para pilha
   - Não cria arrays temporários

3. FATORES QUE AFETAM A PERFORMANCE:

   A) ESCOLHA DO PIVÔ:
   - Primeiro elemento: O(n²) para arrays ordenados
   - Último elemento: O(n²) para arrays ordenados decrescentes
   - Mediana: O(n log n) garantido, mas custo adicional
   - Aleatório: O(n log n) esperado

   B) ESTRATÉGIAS DE OTIMIZAÇÃO:
   - Insertion Sort para arrays pequenos (< 10 elementos)
   - Mediana de três para escolha do pivô
   - 3-way partitioning para arrays com muitos duplicados

4. COMPARAÇÃO COM OUTROS ALGORITMOS:

   Quick Sort vs Merge Sort:
   - Quick: O(n log n) médio, O(n²) pior, O(log n) espaço, não estável
   - Merge: O(n log n) sempre, O(n) espaço, estável

   Quick Sort vs Heap Sort:
   - Quick: O(n log n) médio, O(n²) pior, O(log n) espaço, não estável
   - Heap: O(n log n) sempre, O(1) espaço, não estável

5. ANÁLISE PRÁTICA:

   A) QUANDO USAR QUICK SORT:
   - Arrays grandes com distribuição aleatória
   - Quando espaço é limitado (versão in-place)
   - Quando estabilidade não é importante
   - Implementações otimizadas com pivô inteligente

   B) QUANDO EVITAR QUICK SORT:
   - Arrays pequenos (usar Insertion Sort)
   - Dados parcialmente ordenados
   - Quando estabilidade é crítica
   - Sistemas críticos onde O(n²) é inaceitável

6. VERSÕES OTIMIZADAS:

   A) INTROSORT (Introspecive Sort):
   - Combina Quick Sort + Heap Sort
   - Detecta pior caso e muda para Heap Sort
   - Garante O(n log n) sempre

   B) QUICK SORT COM 3-WAY PARTITIONING:
   - Eficiente para arrays com muitos duplicados
   - Particiona em < pivot, = pivot, > pivot

   C) QUICK SORT HÍBRIDO:
   - Usa Insertion Sort para arrays pequenos
   - Reduz overhead de recursão

O Quick Sort é o algoritmo de ordenação mais usado na prática devido à sua
excelente performance média e implementação in-place, mas requer cuidado
com a escolha do pivô para evitar o pior caso O(n²).
*/
