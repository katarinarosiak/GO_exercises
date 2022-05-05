/*
https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/880/

	i: 32bit int (x)
	o: 32int reversed x

	r:
	- if returned number will be greater that 32 bits => return 0
	- signed means taht it can be negative
	- if signed => return signed
	- if leading zero => drop



	Edge Cases:
	- 0
	-  -123
	-

	- if num > pow(2, 32) => return 0

1.
	take the number
		- coerce to str
		- split into arr of digits
		- reverse
		- if the last is - move to the front

	- if num >  pow(2,31) || num < pow(-2, 23)
		return 0

	2. take the number
		- get the last digit by num % 10 => last digit
		- num / 10
		- lastNum *10 and save
		-

		math.MaxInt32
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// arr := []int{1, 2, 3, 1}
	value := reverse(6927694924)
	fmt.Println(value)
}

func reverse(x int) int {
	lastDig := 0
	reversedNum := 0

	for x != 0 {
		lastDig = x % 10
		x = x / 10

		if reversedNum*10+lastDig > math.MaxInt32 || reversedNum*10+lastDig < math.MinInt32 {
			return 0
		}
		reversedNum = reversedNum*10 + lastDig
	}
	return reversedNum
}
