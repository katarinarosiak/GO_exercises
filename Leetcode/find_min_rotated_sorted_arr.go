package main

import (
	"fmt"
)

func main() {
	a := []int{4, 5, 6, 7, 0, 1, 2}
	value := findMin(a) //2008
	fmt.Println(value)
}

func findMin(nums []int) int {

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) / 2)
		if nums[mid] < nums[mid-1] || len(nums[left:mid]) == 1 {
			return nums[mid]
		} else if allDecreasing(nums[left:mid]) {
			fmt.Println(nums[left:mid])
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -5001
}

func allDecreasing(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}
