package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 3, 0, 6, 1, 5}
	value := missing(arr) //2008
	fmt.Println(value)
}

/*
	i: arr of int  []int{}
		0-N
	o: missing int

		- missing one int
		- not sorted
		- O(N)

		len =
		start =
		iterate through arr of num
		map number num = true

		for loop from start to len +1
			check if num exist in map
			if not return num immidetely
*/

func missing(arr []int) int {
	end := len(arr)
	start := 0
	hash := make(map[int]bool)

	for _, num := range arr {
		hash[num] = true
		if num <= start {
			start = num
		}
	}

	for i := start; i < end; i++ {
		if !hash[i] {
			return i
		}
	}

	return 0
}
