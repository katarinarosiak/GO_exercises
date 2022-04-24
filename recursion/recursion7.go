package main

import (
	"fmt"
)

func main() {
	memo := make(map[int]int)
	value := golamb(8, memo) //
	fmt.Println(value)
}

/*

The following function uses recursion to calculate the Nth number from
a mathematical sequence known as the “Golomb sequence.” It’s terribly
inefficient, though! Use memoization to optimize it. (You don’t have to
actually understand how the Golomb sequence works to do this exercise.)
A:



*/

// def golomb(n)
// return 1 if n == 1
// return 1 + golomb(n - golomb(golomb(n - 1)));
// end

func golamb(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}

	if _, ok := memo[n]; !ok {
		memo[n] = 1 + golamb(n-golamb(golamb(n-1, memo), memo), memo)
	}
	return memo[n]
}

// func addUntil100(arr []int) int {
// 	sum := 0

// 	for _, num := range arr {
// 		if num+sum >= 100 {
// 			return sum
// 		} else {
// 			sum += num
// 		}
// 	}
// 	return sum
// }
