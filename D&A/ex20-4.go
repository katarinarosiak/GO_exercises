package main

import (
	"fmt"
)

func main() {
	arr := []int{5, -10, -6, 9, 4}
	value := higestProf(arr) //2008
	fmt.Println(value)
}

/*
	i: arr of int
	o: count and return int => higest product of any two nums


	- can contain negative nums
	[5, -10, -6, 9, 4] => 60

	higestNeg = 0
	highestPos = 0
	prod = 0


		a = hihgestnegatove * curr
					is prod < a ?
			prod = a
	is it negative?
		Y
			change higestNeg if curr < higestNeg
		N
			change highestPos if curr > highestPos
	return prod


	[5, -10, -6, 9, 4]
*/

func higestProf(arr []int) int {
	highestNeg := 0
	highestPos := 0
	prod := 0
	for _, num := range arr {
		if num < 0 {
			currProd := highestNeg * num
			if currProd > prod {
				prod = currProd
			}

			if highestNeg > num {
				highestNeg = num
			}
		} else {
			currProd := highestPos * num
			if currProd > prod {
				prod = currProd
			}

			if highestPos < num {
				highestPos = num
			}
		}

	}

	return prod
}
