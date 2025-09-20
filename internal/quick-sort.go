package internal

func QuickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1]

	left := []int{}
	right := []int{}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
