/*
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/


i: arr of int  (nums)
o: int => the number of non-duplicate elements in array


R:
sorted increasingly
- remove duplicates (in place)
- each unique el appaers only once
- the same order

- if there there are k el , the first k el of nums should be the result
- then the rest doesn matter

- len of nums >= 1
- len of nyms <= 3 * 10^4
- the element of nums will be between -100 and 100
- sorted non-decreasing

var seen = 0
ar curr = 0
[0, _, _, 1, _, _, 2, _] => [0, 1, 2, _, _, _, _, _]

[0, _, 0, 1, 1, 2, 2]

- iterate through aray
- check if the first number is teh same as seen
	- if yes
		- take the slice (i+1, end)
		- place the one at the end and move all teh other once idx closer

	- if not just continue


	seen 0
	curr 0


*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	value := removeDuplicates(arr)
	fmt.Println(value)
}

// func removeDuplicates(nums []int) int {
// 	hash := make(map[int]bool)
// 	result := len(nums)

// 	for i, num := range nums {
// 		if hash[num] {
// 			nums[i] = 101
// 			result--
// 		} else {
// 			hash[num] = true
// 		}
// 	}
// 	sort.Ints(nums)
// 	return result
// }

/*
seen 4

12340203144
    a
          r


0,0,1,1,1,2,2,3,3,4
a
  r

seen =
a = 0
r = 1

until r < len
is the r seen?
	- if yes
		r++
	- if not:
		seen = r
		swap
		r++
		a++


*/

func removeDuplicates(nums []int) int {

	seen := nums[0]

	a := 0
	r := 0

	for r < len(nums) {
		if nums[r] == seen && r != 0 {
			r++
		} else {
			seen = nums[r]
			nums[a], nums[r] = nums[r], nums[a]
			r++
			a++
		}
	}

	return a
}
