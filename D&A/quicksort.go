package main

import (
	"fmt"
)

func main() {

	value := QuickSort([]int{5, 6, 7, 2, 1, 3}, 0, 5) //true
	fmt.Println(value)
}

/*
A:
array of numbers
leftPointer = first el
rightPointer = len(arr)-2
pivot = len(arr) - 1

iterate arr len
- compare leftpointer to the pivot
- if leftpointer is less than pivot continue
- otherwise stop (return)

do the same with the right pointer

when both pointers has stopped change places

- the left pointer continue
- compare left pointer to pivot
	- if less moves on



- end when the left pointer reached right pointer
	finally swap the left pointer with pivot


	while leftPointer < pivot && rightPointer > pivot

		if leftPointer < pivot && rightPointer > pivot
			if arr[left[] < arr[pivot]
				left++

			if right > pivot
				right--

		else
			temp = arr[leftPointer]
			arr[leftPointer] = arr[rightPointer]
			arr[rightPointer] = temp
	}

	temp = arr[pivot]
	arr[pivot] = arr[left]
	arr[left] = temp

0, 5, 2, 1, 6, 3
*/

// func partition(arr []int) []int {
// 	leftPointer := 0
// 	rightPointer := len(arr) - 2
// 	pivot := len(arr) - 1

// 	for true {
// 		fmt.Println(leftPointer, rightPointer)
// 		for arr[leftPointer] < arr[pivot] {
// 			leftPointer++
// 		}

// 		for arr[rightPointer] > arr[pivot] {
// 			rightPointer--
// 		}

// 		if arr[leftPointer] <= arr[rightPointer] {
// 			fmt.Println("break", arr)
// 			break
// 		} else {
// 			temp := arr[leftPointer]
// 			arr[leftPointer] = arr[rightPointer]
// 			arr[rightPointer] = temp

// 			fmt.Println(arr)
// 			leftPointer++
// 		}
// 	}

// 	// temp := arr[pivot]
// 	// arr[pivot] = arr[leftPointer]
// 	// arr[leftPointer] = temp

// 	return arr
// }

func QuickSort(arr []int, low, high int) []int {
	if low >= high {
		return arr
	}
	pivot := arr[high]
	leftPointer := low
	rightPointer := high

	for leftPointer < rightPointer {

		for arr[leftPointer] <= pivot && leftPointer < rightPointer {
			leftPointer++
		}

		for arr[rightPointer] >= pivot && leftPointer < rightPointer {
			rightPointer--
		}

		Swap(arr, leftPointer, rightPointer)
	}

	Swap(arr, leftPointer, high)

	QuickSort(arr, low, leftPointer-1)
	QuickSort(arr, leftPointer+1, high)
	return arr
}

func Swap(arr []int, idx1 int, idx2 int) {
	temp := arr[idx1]
	arr[idx1] = arr[idx2]
	arr[idx2] = temp
}

// func quickSort(arr []int, low, high int) []int {
// 	if low < high {
// 		var p int
// 		arr, p = partition(arr, low, high)
// 		arr = quickSort(arr, low, p-1)
// 		arr = quickSort(arr, p+1, high)
// 	}
// 	return arr
// }
