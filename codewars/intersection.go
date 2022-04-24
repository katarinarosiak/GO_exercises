package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{0, 2, 4, 6}
	value := intersection(arr1, arr2) //[2,4]
	fmt.Println(value)
}

/*
func => returns intersection of two arrs
- intersation is a third array contained within both of the arrays
- O(n)


	A:
	- find longer arr
	- iterate through the longer array
	- create a hash of the elements, use them as keys and true as a value in a hash table
	- iterate through second arr
	- if el is in the hash push to new slice
	- return slice
*/
//1, 2, 3, 4, 5
func intersection(arr1 []int, arr2 []int) []int {
	var inter []int
	hashTable := make(map[int]bool)
	var longer []int
	var shorter []int

	if len(arr1) > len(arr2) {
		longer = arr1
		shorter = arr2
	} else {
		longer = arr2
		shorter = arr1
	}

	for _, el := range longer {
		hashTable[el] = true
	}

	for _, num := range shorter {
		if hashTable[num] {
			inter = append(inter, num)
		}
	}

	return inter
}
