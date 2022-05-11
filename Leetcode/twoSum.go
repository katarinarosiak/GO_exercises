/*
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice.
You can return the answer in any order.

i:array of int (nums)
- int (target)

o:
indexes of two numbers that adds up to target


r:
- each input have exactly one solution
- don\t use the same element twice
- answer in any order
- the length of nums >= 2 and <= 4 pow of 10
- indexes will be -10 pow 9 <= nums[i] <= 10 pow 9
- targt: -10 pow 9 <= target <= 10 pow 9


[2,7,11,15]   9
O(N pow 2)   => O(N)

2
[11,-15, 2,22, 7, 44]

Test Casas:
[3, 3], 6 => [0, 1]
[2, -1, 11, 15 -4], 5 => [1, 4]
[1, 6, 5, 3,  0] , 5 => [2, 4]


{
	3:true
}

[3, 3]


- map the numbers from the array into a hash {1:true}
- iterate through array and for each number check if there
is a number that will add up to the target num
- find the index of that number in the array
- if the indexes are not the same
	- return that and the current array

*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 6, 11, -1, 3, 0}
	value := twoSum(arr, 6) //[2, 4]]
	fmt.Println(value)
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if !hash[nums[i]] {
			hash[nums[i]] = true
		}
	}

	for i := 0; i < len(nums); i++ {
		searchedNum := target - nums[i]
		if hash[searchedNum] {
			idx1 := indexOf(searchedNum, nums)

			if idx1 != i {
				return []int{idx1, i}
			}
		}
	}
	return nil
}

func indexOf(num int, arr []int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
}

// func twoSum(nums []int, target int) []int {
// 	seen := make(map[int]int, 2)
// 	indeces := make([]int, 2)
// 	var complement int

// 	for idx, num := range nums {
// 			complement = target - num
// 			if complementIdx, ok := seen[complement]; ok {
// 					indeces[0] = complementIdx
// 					indeces[1] = idx
// 					break
// 			} else {
// 					seen[num] = idx
// 			}
// 	}
// 	return indeces
// }

// func twoSum(nums []int, target int) []int {
// 	m := make(map[int]int)
// 	output := make([]int, 2)

// 	for idx1, val1 := range nums {
// 			delta := target - val1
// 			idx2, present := m[delta]

// 			if present {
// 					output[0], output[1] = idx1, idx2
// 					break
// 			} else {
// 					m[val1] = idx1
// 			}
// 	}
// 	return output
// }
