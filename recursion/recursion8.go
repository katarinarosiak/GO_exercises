package main

import (
	"fmt"
)

func main() {
	memo := make(map[[2]int]int)
	value := uniquePaths(3, 3, memo) //
	fmt.Println(value)
}

/*

Here is a solution to the “Unique Paths” problem from an exercise in the
previous chapter. Use memoization to improve its efficiency:

*/

func uniquePaths(rows int, cols int, memo map[[2]int]int) int {
	if rows == 1 || cols == 1 {
		return 1
	}

	arr := [2]int{rows, cols}
	if _, ok := memo[arr]; !ok {
		memo[arr] = uniquePaths(rows-1, cols, memo) + uniquePaths(rows, cols-1, memo)
	} else {
		return memo[arr]
	}

	// return uniquePaths(rows - 1, col, memo) + uniquePaths(rows, col - 1, memo)
}
