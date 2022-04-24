package main

import (
	"fmt"
)

func main() {

	value := indexOf("abcccx", 0) //23
	fmt.Println(value)
}

/*

This problem is known as the “Unique Paths” problem: Let’s say you have
a grid of rows and columns. Write a function that accepts a number of rows
and a number of columns, and calculates the number of possible “shortest”
paths from the upper-leftmost square to the lower-rightmost square.
For example, here’s what the grid looks like with three rows and seven
columns. You want to get from the “S” (Start) to the “F” (Finish).


func => calc. the number of possible shortest paths from left upper corner
  tp right-lower corner
i = number (rows), number (columns)
o = number of possible patsh

A:


	if s == ""
		return -1

	if s[0] == x
		return idx
	else indexOf(s[currIdx:], idx+1)
*/

func indexOf(s string, idx int) int {

	if len(s) == 0 {
		return -1
	}

	if s[0] == 'x' {
		return idx
	} else {
		return indexOf(s[1:], idx+1)
	}
}
