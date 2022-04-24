package main

import (
	"fmt"
)

func main() {

	value := countChar([]string{"ab", "c", "def", "ghij", "i"}) //10
	fmt.Println(value)
}

/*

Use recursion to write a function that accepts an array of strings and
returns the total number of characters across all the strings. For example,
if the input array is ["ab", "c", "def", "ghij"] , the output should be 10 since there
are 10 characters in total.


func =>
i = array of str
o = total num of chars from all teh strings

	A:

*/

func countChar(words []string) int {

	if len(words) == 0 {
		return 0
	}
	return len(words[0]) + countChar(words[1:])
}
