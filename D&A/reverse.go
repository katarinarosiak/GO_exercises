package main

import (
	"fmt"
)

func main() {

	value := reverse([]string{"abode", "ABc", "xyzD"}) //{4,3,1}
	fmt.Println(value)
}

// Create a new function to reverse an array that takes up just O(1) extra space.
func reverse(arr []string) []string {
	len := len(arr)
	for i := len - 1; i >= 0; i-- {
		arr = append(arr, arr[i])
	}
	return arr[len:]
}
