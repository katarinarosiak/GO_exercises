package main

import (
	"fmt"
)

func main() {
	var a1 = []int{5, 25, 75}
	// var a2 = []string{"bcc", "bccc", "lade", "lade", "abada"}

	value := twoSum(a1, 100)
	fmt.Println(value)

}

/*
https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/submissions/
i: arry of int (sorted) - non-decreasing
	int = terget

o: int, indexes


R:
- two num that adds up to target
- index1 >= 1
- index1 <= index2
index2 <= arr len

- only one solution
- use eahc element only once
- use only constant extra space


[5 25 75] , 9
s
       e

[1 2 3 4 5] 7
 s
         e
set
start 0
end = lats el

take two nums,
	if they sum up to target return indexes
	if not check if sum > target
		end++
	if sum < target
		start++

	return -1
*/

func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1

	for start < end {
		sum := numbers[start] + numbers[end]
		if sum == target {
			return []int{start + 1, end + 1}
		} else if sum > target {
			end--
		} else {
			start++
		}
	}

	return nil
}
