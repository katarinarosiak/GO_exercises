package main

import (
	"fmt"
)

func main() {

	value := Factorial(5) //
	fmt.Println(value)
}

/*

!5 = 5*4*3*2*1
n - Factorial(n-1)
*/

func Factorial(n int) int {
	if n < 1 {
		return 1
	}
	return n * Factorial(n-1)
}
