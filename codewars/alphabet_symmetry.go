package main

import (
	"fmt"
	"strings"
)

func main() {

	value := solve([]string{"abode", "ABc", "xyzD"}) //{4,3,1}
	fmt.Println(value)
}

/*
i: arr of words
o: array of numers of letter that occupy
their position in the alphabeth for each word
solve(["abode","ABc","xyzD"]) = [4, 3, 1]

- no spaces
- upper and lower case
- case insensitive

abcdefghi j  k  l  m  n  o  p  q  r  s  t  u  v  w  x  y  z
123456789 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26

create alpha string
[]

- ietaret through the given arr
	var count
	- for each word ietaret through each letter
		if the letter == alphastr[idx]
		count++

	append count to empty arr


*/

func solve(slice []string) []int {
	alphaStr := "abcdefghijklmnopqrstuvwxyz"
	final := make([]int, len(slice))

	for sliceIdx, word := range slice {
		fmt.Println(word)
		for idx, char := range strings.ToLower(word) {
			if string(char) == string(alphaStr[idx]) {
				fmt.Println(int(char), string(char), idx+1)
				final[sliceIdx]++
			}
		}

	}
	return final
}
