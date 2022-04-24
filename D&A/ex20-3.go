package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 7, 5, 8, 11, 2, 6}
	value := profit(arr) //2008
	fmt.Println(value)
}

/*
	i: arr of int


	[1,4,1,6,7,11] = 30   1, 11 => 10
	[7,8,4,7, 1,6] = 33
	[0,11,5,2,9,4] = 31
	[10, 7, 5, 8, 11, 2, 6]

	profit = 6    if higehr
	lowest = 2    if lower


	take the fild el as lowest
	profit = 0

	iterate through the arr
	on each iteration take two numbers
	start from idx 1
	curr

	is there a profit?
		curr - lowest > 0 ?
		if yes:
		profit = curr - lowest

	is curr < lowest ?
		if yes
		lowest = curr



	return profit

		[10, 7, 5, 8, 11, 2, 6]
*/

func profit(arr []int) int {
	lowest := arr[0]
	profit := 0

	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		fmt.Println(curr, "curr")
		fmt.Println(lowest, "lowest")
		fmt.Println(profit, "profit")

		currProf := curr - lowest
		if currProf > 0 && currProf > profit {
			profit = currProf
		}

		if curr < lowest {
			lowest = curr
		}
		fmt.Println(profit, "now")
	}

	return profit
}
