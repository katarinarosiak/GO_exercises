package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var a1 = []string{"aba", "ccc", "ade", "ccc"}
	var a2 = []string{"bcc", "bccc", "lade", "lade", "abada"}

	value := InArray(a1, a2)
	fmt.Println(value)

}

/*
i: arr of str, arr of str

o: sorted array

- sorted: \
- lexicographical order of the strings of a1
which are substrings of strings of a2
- r (no diplicates)
- if [] => []

[aba, ccc, ade, ccc]
[bcc, bccc, lade, lade, abada]
=> ccc, ade, aba
=> aba ade ccc

aba
ccc
ade

- join a2 into a string with ' '
- itearte through a1
- if the current el contains in a string
	- and slice doesn't include yet
	- append to slice
- sort the slice
- return

*/

func InArray(array1 []string, array2 []string) []string {
	final := []string{}
	str := strings.Join(array2, " ")
	hash := make(map[string]bool)

	for _, substr := range array1 {
		if strings.Contains(str, substr) {
			if _, value := hash[substr]; !value {
				hash[substr] = true
				final = append(final, substr)
			}
		}
	}

	sort.Strings(final)

	return final
}
