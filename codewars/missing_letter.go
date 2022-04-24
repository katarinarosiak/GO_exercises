package main

import (
	"fmt"
)

func main() {

	value := missingLett("the quick brown box jumps over the lazy dog") //f
	fmt.Println(value)
}

/*
func => uses a stack to reverse a string
i = string (all alpha except 1)
o = return missing letter

	A:
	- create a hash map from the string
	- iterate through arr of all alpha chars
	- if alpha doesn't exist in teh hash return that el

*/

func missingLett(str string) string {
	letters := make(map[byte]bool)
	alpha := "abcdefghijklmnopqrstuvxyz"

	for idx := 0; idx < len(str); idx++ {
		letters[str[idx]] = true
	}

	for idx := 0; idx < len(alpha); idx++ {
		if !letters[alpha[idx]] {
			return string(alpha[idx])
		}
	}
	return ""
}
