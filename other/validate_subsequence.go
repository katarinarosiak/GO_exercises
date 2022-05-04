/*
https://www.algoexpert.io/questions/Validate%20Subsequence

i: arr of nym  [], arr o num
o: bool



R;
- numbers: integers, can be negative,
- don't have be to adjacent but are in the same order
- the numbers n arr1 has to exist in array 2 in the right order

- what if []
- what if not the same order
- the secod is greater


i1 : 3    [1, 2, 3]
i2 : 1    [3, 2, 1]


	if i2 >= arr2 len => return true
	if i1 >= arr1 len => return false

- iterate through the arr1 (longer)
- if arr1[i1]=== arr2[i2]
	if not:
		- increment i1
	if true
		- increment i1, i2

*/
package main

import (
	"fmt"
)

func main() {
	arr1 := []int{0, 1, 4, 2, 3}
	arr2 := []int{4, 1, 0}
	value := IsValidSubsequence(arr1, arr2) //true
	fmt.Println(value)
}

func IsValidSubsequence(array []int, sequence []int) bool {
	idx2 := 0

	if len(sequence) == 0 || len(array) == 0 || len(sequence) > len(array) {
		return false
	}

	for _, num := range array {
		if num == sequence[idx2] {
			idx2++
		}
		if idx2 >= len(sequence) {
			return true
		}
	}
	return false
}
