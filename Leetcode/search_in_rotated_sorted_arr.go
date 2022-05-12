package main

import (
	"fmt"
)

func main() {
	a := []int{3, 4, 5, 6, 0, 1, 2}
	value := search(a, 0) //2008
	fmt.Println(value)
}

func search(nums []int, target int) int {

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) / 2)

		if nums[mid] == target {
			return mid
		} else if nums[left] == target {
			return left
		} else if nums[right] == target {
			return right
		} else if nums[left] <= nums[mid] && nums[left] <= target && target <= nums[mid] ||
			nums[mid] <= nums[right] && (nums[mid] <= target && target <= nums[right]) != true {
			right = mid - 1
		} else if nums[mid] <= nums[right] && nums[mid] <= target && target <= nums[right] ||
			nums[left] <= nums[mid] && (nums[left] <= target && target <= nums[mid]) != true {

			left = mid + 1
		}
	}
	return -1
}
