package main

import (
	"fmt"
)

func main() {

	value := indexOf("abcccx", 0) //23
	fmt.Println(value)
}

/*

Use recursion to write a function that accepts a string and returns the
first index that contains the character “x.” For example, the string,
"abcdefghijklmnopqrstuvwxyz" has an “x” at index 23. To keep things simple,
assume the string definitely has at least one “x.”


func =>
i = str
o = first idx that contains char x (int)

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
