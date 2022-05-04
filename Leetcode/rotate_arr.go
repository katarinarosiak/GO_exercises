/*
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/646/

i: arr of int, k = int
o: the same array


r: k is non negative
- arr len is at lesat 1 and can be 10^5
- can be big numbers
k >= 0
k <= 10^5

a = last

iterate start from end
	arr[i] = arr[i-1]
	ar[0] = last
[1,2, 3, 4]

iterate k times
	- rotate()
		- take teh first el shift (save in variable)
		- append to the end of the array



[1,2,3] ,3

[1, 2, 3]



*/
package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6}
	value := rotate(arr, 2)
	fmt.Println(value)
}

// [ 5, 6, 0, 1, 2, 3, 4]
//

// func rotate(nums []int, k int) []int {
// 	newArr := make([]int, len(nums))

// 	for i, num := range nums {

// 		if i+k >= len(nums) {
// 			idx := (i + k) % len(nums)
// 			newArr[idx] = num
// 		} else {
// 			newArr[i+k] = num
// 		}
// 	}

// 	for i, num := range newArr {
// 		nums[i] = num
// 	}

// 	return nums
// }

func rotate(nums []int, k int) []int {
	k = k % len(nums)
	left, right := 0, len(nums)-1

	reverse(nums, left, right)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)

	return nums
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left, right = left+1, right-1
	}
}
