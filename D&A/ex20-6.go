/*
i: arr of unsorted int
r: len of longest consecutive sequence


S;
- int that increase by 1

- in array don't have to be adjacent
- do it in O(N)

{
10
5
12
3
}

max 3
curr 1
5

10


[10, 5, 12, 3, 55, 30, 4, 11, 2]   => 2-3-4-5



*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3}
	value := longestSequenceLen(arr)
	fmt.Println(value)
}

func longestSequenceLen(arr []int) int {
	hash := make(map[int]bool)
	maxLen := 0

	for _, num := range arr {
		if !hash[num] {
			hash[num] = true
		}
	}

	for _, num := range arr {
		if !hash[num-1] {

			currLen := 1
			currNum := num

			for hash[currNum+1] {
				currNum++

				currLen++

				if currLen > maxLen {
					maxLen = currLen
				}
			}

		}
	}
	return maxLen
}
