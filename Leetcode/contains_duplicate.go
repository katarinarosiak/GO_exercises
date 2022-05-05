/*
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/578/
i: int arr (nums)
o: bool

- return tru eof any value appears at least twice
- false otherwise







*/
package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 1}
	value := containsDuplicate(arr)
	fmt.Println(value)
}

func containsDuplicate(nums []int) bool {
	hash := make(map[int]bool)

	for _, num := range nums {
		if hash[num] {
			return true
		} else {
			hash[num] = true
		}
	}
	return false
}
