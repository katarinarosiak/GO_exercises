package main

import (
	"fmt"
)

func main() {
	var a1 = []int{2, 7, 11, 15}
	// var a2 = []string{"bcc", "bccc", "lade", "lade", "abada"}

	value := twoSum(a1, 9)
	fmt.Println(value)

}

/*
13.04
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


[2,6,11,12] , 9
 s
e


set
start 0
end = lats el

take two nums,
	if they sum up to target return indexes
	if not check if sume is > target
		decrease the end


	return -1
*/

func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1

	for start < end {
		sum := numbers[start] + numbers[end]
		if sum == target {
			return []int{start, end}
		} else {
			if sum > target {
				return []int{}
			}
		}
	}
	return []int{}
}
