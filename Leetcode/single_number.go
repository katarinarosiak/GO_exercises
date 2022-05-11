/*
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/549/

i: arr of int




r; arr len at least 1
- every el appear twice except for 1
- find taht el

seen = 2

[2, 1, 1, 5, 4 ,2, 4 ]

*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 1, 1, 5, 4, 2, 4}
	value := singleNumber(arr)
	fmt.Println(value)
}

//solution 1
// func singleNumber(nums []int) int {
// 	hash := make(map[int]int)

// 	for _, num := range nums {
// 		if hash[num] == 0 {
// 			hash[num] = 1
// 		} else {
// 			hash[num]++
// 		}
// 	}

// 	for _, num := range nums {
// 		if hash[num] == 1 {
// 			return num
// 		}
// 	}
// 	return -1
// }

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res = n ^ res
	}
	return res
}
