package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	a1 := []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	a2 := []string{"bbbaaayddqbbrrrv"}

	value := test(a1, a2) //edcba
	fmt.Println(value)
}

/*
You are given two arrays a1 and a2 of strings.
Each string is composed with letters from a to z.
Let x be any string in the first array and y be any string in the second array.

Find max(abs(length(x) − length(y)))

If a1 and/or a2 are empty return -1 in each language except in Haskell (F#)
where you will return Nothing (None).

i: 2 arr of strings: a1, a2



- each string: letters from a-z
- knowing that:
- x => any string in teh first arr
- y => any string in the second arr

find max abs value of pair of len of two arrays
 lenght of x and length of y

- if a1 or a2 are empty return -1
- x => first arr
- y => sec arr

- sort by length both arrays
-

[a a ab abb aaaa aaaaa]
[abcd, aaa]

- join arr
- sort by len
- take the first and the last deduct
- count abs
return abs

"cccooommaaqqoxii" 16   3
*/

func test(a1, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	sort.Slice(a1, func(i, j int) bool {
		return len(a1[i]) < len(a1[j])
	})
	sort.Slice(a2, func(i, j int) bool {
		return len(a2[i]) < len(a2[j])
	})

	firstA1 := len(a1[0])
	secondA1 := len(a1[len(a1)-1])

	firstA2 := len(a2[0])
	secondA2 := len(a2[len(a2)-1])

	abs1 := int(math.Abs(float64(firstA1 - secondA2)))
	abs2 := int(math.Abs(float64(firstA2 - secondA1)))

	if abs1 > abs2 {
		return abs1
	} else {
		return abs2
	}
}
