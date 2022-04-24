package main

import (
	"fmt"
)

func main() {

	value := triangular(3) //28
	fmt.Println(value)
}

/*

There is a numerical sequence known as “Triangular Numbers.” The
pattern begins as 1, 3, 6, 10, 15, 21, and continues onward with the Nth
number in the pattern, which is N plus the previous number. For example,
the 7th number in the sequence is 28, since it’s 7 (which is N) plus 21
(the previous number in the sequence). Write a function that accepts a
report erratum • discussChapter 11. Learning to Write in Recursive
number for N and returns the correct number from the series. That is, if
the function was passed the number 7, the function would return 28.


func =>
i = number n
o = int from the series

traingular numbers:
1, 3, 6, 10, 15, 21

N + pervious num
1+0 2+1, 3+3, 4+6, 5+10


A:
	- 7 + trianglaur(6)




*/

func triangular(num int) int {
	if num == 0 {
		return 0
	}
	return num + triangular(num-1)
}
