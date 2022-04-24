package main

import (
	"fmt"
	"sort"
)

func main() {

	value := test([]int{6, 5, 3, 2, 1}) //edcba
	fmt.Println(value)
}

/*
The following function finds the “missing number” from an array of inte-
gers. That is, the array is expected to have all integers from 0 up to the
array’s length, but one is missing. As examples, the array, [5, 2, 4, 1, 0] is
missing the number 3, and the array, [9, 3, 2, 5, 6, 7, 1, 0, 4] is missing the
number 8.
Here’s an implementation that is O(N 2 ) (the includes method alone is already
O(N), since the computer needs to search the entire array to find n ):
Use sorting to write a new implementation of this function that only takes
O(N log N). (There are even faster implementations, but we’re focusing on
using sorting as a technique to make code faster.)


- i: array of positiv int
- o: int => int => missing number

- use sorting
- O (N log N)

	A:
  - sort arr
	- iterate through arr every other
	- take two numbers and
		if i+1 > len of arr
			previos != current +1
			return current
		else
			if current != next +1
				return
	return -1
	[1,2,3,5,6]
*/

func test(arr []int) int {
	sort.Ints(arr)

	for i, num := range arr {
		if i+1 == len(arr) {
			if num-1 != arr[i-1] {
				return num - 1
			}
		} else {
			if num+1 != arr[i+1] {
				return num + 1
			}
		}
	}
	return -1
}
