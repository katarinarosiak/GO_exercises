package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	value := Revrot("123456987654", 6) //"234561 876549" 654321 965432
	fmt.Println(value)
}

/*
i: str of digit
i: int n
o: string

r:
- ignore the last chunk if rhe size is less than n
- cut the str into chunks => size == n
- if chunk represnets int:
		 suc as the sum of the cubes of its digits is divisible by 2
		 => reverse the chunk
		 else
		 roate to the left by one postion

	put together chunks and return as a string


	if n <= 0 or str == "" => ""
	if n > than the len of str => return ""

	"123456 987653", 6) --> "234561 356789"
	final :=

	take str len % n  == 0
		numOfSlices = str len % n
	else
		numOfSlices - 1

	iteare numOfSlices times idx  idx+ numOfSlices
		on each ietration make a slice
			from [idx: idx+  numOfSlices]

			if calcul(slice)
				rotate
				append to final
			else
				reverse
				append to final

	return final joined to string

	calcul:
		[d, d, d, d]
		iterate trhough the slice
			take the str convert to int and cube it num*num
			sum the results
			check if divisible by 2
			return true else false

	rotate
		take the first el of the slice
		plcae it in the end
		return the slice
	reverse
		[]string
		iterate through the slice
		append in the beginning of a slice



		iteare slicenUm times
		123456   3

		strat := 0
		end := start+n
		start+ n
		0:3  3:6
*/

func Revrot(s string, n int) string {
	if n <= 0 || len(s) == 0 || n > len(s) {
		return ""
	}

	numOfSlices := len(s) / n
	final := []string{}
	start := 0
	end := start + n

	for i := 0; i < numOfSlices; i++ {
		sl := s[start:end]
		start += n
		end += n
		if calcul(sl) {
			final = append(final, reverse(sl))
		} else {
			final = append(final, rotate(sl))
		}
	}

	return strings.Join(final, "")
}

func calcul(sl string) bool {
	result := 0
	for _, num := range sl {
		conv, _ := strconv.Atoi(string(num))
		cube := conv * conv * conv
		result += cube
	}
	return result%2 == 0
}

func rotate(sl string) string {
	return sl[1:] + sl[0:1]
}

func reverse(sl string) string {
	str := ""
	for _, char := range sl {
		str = string(char) + str
	}
	return str
}
