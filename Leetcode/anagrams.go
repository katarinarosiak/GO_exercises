package main

import (
	"fmt"
)

func main() {

	value := anagrams("anagram", "margana") //2008
	fmt.Println(value)
}

func anagrams(s string, b string) bool {
	hash := make(map[rune]int)

	for _, char := range s {
		if hash[char] == 0 {
			hash[char] = 1
		} else {
			hash[char]++
		}
	}
	
	for _, char 

	return true
}
