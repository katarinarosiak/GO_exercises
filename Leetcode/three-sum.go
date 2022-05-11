//not finished

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	value := threeSum(arr)
	fmt.Println(value)
}

func threeSum(nums []int) [][]int {
	pointer := 0
	start := 1
	end := len(nums) - 1
	final := [][]int{}

	sort.Ints(nums)
	pointerSeen := 99999

	for pointer > 0 {
		if pointerSeen == nums[pointer] {
			continue
		}
		fmt.Println(11)
		for start < end {
			startSeen := 999999
			target := nums[pointer] * -1
			sum := nums[start] + nums[end]

			if sum < target {
				start++
			} else if sum > target {
				end--
			} else if sum == target && startSeen != nums[start] {
				final = append(final, []int{nums[pointer], nums[start], nums[end]})
				startSeen = nums[start]
				start++
			}
		}

		pointerSeen = nums[pointer]
		pointer++
		start = pointer + 1
		end = len(nums) - 1
	}

	return final
}

/*
{

[-1, -1, 1] [-1 -1 1] [-1, 0, 1]

[-4 -1,-1 -1 0,0, 1,1, 2]
  p
     s
			                 e

8
pointer = 0
start = 1
end = len(arr)-1
arr = []

seenS = {
 -1
}

sort arr
itertae until pointer == len(arr)-3
	for each iteration:
		for s < e
				get the curr*-1

				if sum < target
					s++
				else if sum > target
				  e--
				else if sum == target and s not in seen
					arr append arr[p], arr[s], arr[e]
					add to seen (s)
					s++
		p++
		s=p+1
		e = end
*/

// func twoSum(nums []int, target int) [][]int {
// 	start := 0
// 	end := len(nums) - 1
// 	twos := [][]int{}

// 	for start < end {
// 		if (nums[start] + nums[end]) < target {
// 			start++
// 		} else if nums[start]+nums[end] > target {
// 			end--
// 		} else {
// 			twos = append(twos, []int{nums[start], nums[end]})
// 			start++
// 		}
// 		for start < end && start != 0 && nums[start] == nums[start-1] {
// 			start++
// 		}
// 	}
// 	return twos
// }

// func threeSum(nums []int) [][]int {
// 	sort.Ints(nums)
// 	triplets := [][]int{}

// 	if len(nums) < 3 {
// 		return triplets
// 	}

// 	var complements [][]int
// 	var triplet []int
// 	anchor := 0

// 	for anchor < len(nums)-2 {
// 		for anchor < len(nums)-2 && anchor != 0 && nums[anchor] == nums[anchor-1] {
// 			anchor++
// 		}
// 		complements = twoSum(nums[anchor+1:], -nums[anchor])
// 		for _, twos := range complements {
// 			triplet = append(twos, nums[anchor])
// 			triplets = append(triplets, triplet)
// 		}
// 		anchor++
// 	}
// 	return triplets
// }
