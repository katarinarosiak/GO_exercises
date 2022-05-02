package main

import (
	"fmt"
)

func main() {
	var a1 = []int{2, 4, 0, 100, 4, 11, 2602, 36}
	// var a2 = []string{"bcc", "bccc", "lade", "lade", "abada"}

	value := FindOutlier(a1)
	fmt.Println(value)

}

/*
i: array if int



- len >= 3
- can be very large
- either even or odd, except for single one
- 0 is even

	if le arr == 1
		return arr[0]

	if sum(slice1) % 2 != 0
		return f(slice1)

	else
		return f(slice2)


*/

func FindOutlier(integers []int) int {
	searchOdd := determineSearchOdd(integers[0:3])

	for _, num := range integers {
		if searchOdd && num%2 != 0 {
			return num
		} else if !searchOdd && num%2 == 0 {
			return num
		}
	}
	return 0
}

func determineSearchOdd(slice []int) bool {
	odd := 0
	even := 0

	for _, num := range slice {
		if num%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	return even > odd
}
