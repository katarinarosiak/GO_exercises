package main

import (
	"fmt"
	"sort"
)

func main() {

	value := test([]int{1, 5, 7, 9, 11, 45, 1, 4, 121}) //edcba
	fmt.Println(value)
}

/*
Given an array of positive numbers, write a function that returns the
greatest product of any three numbers. The approach of using three
nested loops would clock in at O(N 3 ), which is very slow. Use sorting to
implement the function in a way that it computes at O(N log N) speed.
(There are even faster implementations, but we’re focusing on using
sorting as a technique to make code faster.)


- i: array of positiv int
- o: int => greatest product of any three numbers

- use sorting
- O (N log N)

	A:
  - sort
	- choose three last and * them
*/

func test(arr []int) int {
	sort.Ints(arr)
	nums := arr[len(arr)-3:]
	prod := 1

	for _, i := range nums {
		prod *= i
	}

	return prod
}
