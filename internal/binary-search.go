package internal

import "fmt"

func BinarySearch(arr []int, target int) (int, error) {

	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid, nil
		}

		// se o elemento do meio for menor que o target, o target está na metade direita
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1, fmt.Errorf("target not found")
}
