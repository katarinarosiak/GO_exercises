package main

import (
	"fmt"
)

func main() {

	value := addUntil100([]int{1, 2, 40, 20, 70}) //63
	fmt.Println(value)
}

/*

The following function accepts an array of numbers and returns the sum,
as long as a particular number doesn’t bring the sum above 100. If adding
a particular number will make the sum higher than 100, that number is
ignored. However, this function makes unnecessary recursive calls. Fix
the code to eliminate the unnecessary recursion:

A:



*/

// func addUntil100(arr []int{}]) int {

// 	if len(arr) == 0 {
// 		return 0
// 	}

// 	if arr[0] + addUntil100(arr[1:len(arr)-1]) > 100 {
// 		return addUntil100(arr[1 : len(arr) - 1])
// 	} else {
// 		return arr[0] + addUntil100(arr[1 : len(arr)])
// 	}
// }

func addUntil100(arr []int) int {
	sum := 0

	for _, num := range arr {
		if num+sum >= 100 {
			return sum
		} else {
			sum += num
		}
	}
	return sum
}
