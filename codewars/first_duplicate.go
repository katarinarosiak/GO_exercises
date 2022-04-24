package main

import (
	"fmt"
)

func main() {
	arr := []string{"a", "b", "c", "d", "c"}
	value := firstDupl(arr) //c
	fmt.Println(value)
}

/*
func => accepts arr of strings
returns first duplicate value it finds

- efficiency O(n)
- only one duplicate in array

	A:
	- iterate throgh array
	- map the arr of strings
	- if el already exist return the el and breakj

*/
//1, 2, 3, 4, 5
func firstDupl(arr []string) string {
	hashTable := make(map[string]bool)

	for _, el := range arr {
		if hashTable[el] {
			return el
		} else {
			hashTable[el] = true
		}
	}
	return ""
}
