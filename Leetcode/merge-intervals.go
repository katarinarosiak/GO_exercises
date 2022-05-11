//https://leetcode.com/problems/merge-intervals/submissions/
/*
- if len 0 or 1 just return

- []
- iterate until the end of arr
- if array not empty:
	- take last el and current
	- check if they overlap:
		- if yes:
			- merge them and replace in the array
		- if no:
			- push the el from array to final
			- empty array
			- push the current to empty arr
 - if arr empty push current

- push the arr el to final
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [][]int{{1, 4}, {2, 3}}
	value := merge(a)
	fmt.Println(value)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.SliceStable(intervals[:], func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	final := make([][]int, 0)
	temp := make([][]int, 0)

	for _, curr := range intervals {
		if len(temp) == 0 {
			temp = append(temp, curr)
		} else {
			last := temp[len(temp)-1]
			if isOverlap(last, curr) {
				temp[0] = combine(last, curr)
			} else {
				final = append(final, last)
				temp = make([][]int, 0)
				temp = append(temp, curr)
			}
		}
	}
	final = append(final, temp[0])
	return final
}

func combine(arr1 []int, arr2 []int) []int {
	arr1 = append(arr1, arr2...)
	var min int
	var max int

	if arr1[0] > arr2[0] {
		min = arr2[0]
	} else {
		min = arr1[0]
	}

	if arr1[1] > arr2[1] {
		max = arr1[1]
	} else {
		max = arr2[1]
	}

	return []int{min, max}
}

func isOverlap(arr1 []int, arr2 []int) bool {
	el := arr2[0]
	return arr1[0] <= el && arr1[1] >= el
}
