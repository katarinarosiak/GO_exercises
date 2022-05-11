package main

import (
	"fmt"
)

func main() {
	var a1 = []int{0, 1, 2, 3, 4}
	// var a2 = []string{"bcc", "bccc", "lade", "lade", "abada"}

	value := binarySearch(a1, 1)
	fmt.Println(value)

}

// with recursion:
// func binarySearch(arr []int, target int) int {

// 	if len(arr) == 0 {
// 		return -1
// 	}
// 	current := arr[len(arr)/2]
// 	fmt.Println(current, "curr")
// 	if current == target {
// 		return indexOf(current, arr)
// 	} else if target > current {
// 		return binarySearch(arr[len(arr)/2:], target)
// 	} else {
// 		return binarySearch(arr[:len(arr)/2], target)
// 	}
// }

// func indexOf(el int, arr []int) int {
// 	for i, num := range arr {
// 		if num == el {
// 			return i
// 		}
// 	}
// 	return -1
// }

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right { // <= here because left and right could point to the same element, < would miss it
		mid := left + (right-left)/2 // use `(right - left) /2` to prevent `left + right` potential overflow
		// found target, return its index
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			// middle less than target, discard left half by making left search boundary `mid + 1`
			left = mid + 1
		} else {
			// middle greater than target, discard right half by making right search boundary `mid - 1`
			right = mid - 1
		}
	}
	return -1 // if we get here we didn't hit above return so we didn't find target
}
