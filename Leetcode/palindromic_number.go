/*
https://leetcode.com/problems/palindrome-number/

i: integer x
o: boolean


- check if an integer is a palindormic int:
	- reads the same backwards and forwards
	- 121 - true
	- 123 - false

	- input can be negative => false



	123

	1 * 1
	+
	2 * 10
	+
	3 * 100


	reverse the number
	- get the digits separetly
	- add them backwards
	=> reversed number


	[3, 2, 1]

	div := 10
	rest := 0
	final := 321
	- take the number and
		rest = num % div
		num - rest
		div *10
		res / div
		push to arr

		len-1 = 2
		- iterate 10pow2 i /10
		final += arr[i] * i


	div:= 10
	rest:= 0
	final:=0
	arr := []

	iterate until num % div == num
		rest = num % div
		num -= rest
		div *= 10
		push to array (res/div)  => first digit

	- arrlen = check the length of arr
	iterate: start = 10 pow arrlen; end = 0   i /= 10
		final += arr[i]*i

*/

package main

import (
	"fmt"
)

func main() {

	value := isPalindrome(1001) //false
	fmt.Println(value)
}

func isPalindrome(x int) bool {

	if x < 0 || x%100 == 0 && x != 0 {
		return false
	}

	revertedNum := 0

	for x > revertedNum {
		revertedNum = revertedNum*10 + x%10
		x /= 10
	}

	// div := 10
	// rest := 0
	// final := 0
	// arr := []int{}
	// num := x

	// for num%div != 0 {
	// 	rest = num % div
	// 	num = num / 10
	// 	fmt.Println(num)
	// 	arr = append(arr, rest/(div/10))
	// 	div *= 10
	// }
	// fmt.Println(arr)
	// arrlen := float64(len(arr) - 1)
	// pow := math.Pow(10, arrlen)

	// for _, num := range arr {
	// 	final += num * int(pow)
	// 	pow = pow / 10
	// }

	// return x == final
}
