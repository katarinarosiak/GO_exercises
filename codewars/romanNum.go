package main

import (
	"fmt"
)

func main() {

	value := Decode("MMVIII") //2008
	fmt.Println(value)
}

/*
DXLVII - 547
MCDLXXX - 1480

500
40
5
1
1

iterate through string
	if the slice is only one
		sum += el
take a slice of len 2
if first < second
	sum += second - first
	i +2
else
	sum += first+second


i: roman numeral
r: value as numeric decimal int

R:
- 1000 = M

I
II
III
IV
V
VI
VII
VIII
IX


IV

2008
MM VI II


*/

func Decode(roman string) int {
	hash := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	for i := 0; i < len(roman); i++ {
		if a := roman[i:]; len(a) == 1 {
			sum += hash[rune(a[0])]
		} else {
			chars := roman[i : i+2]
			first := hash[rune(chars[0])]
			second := hash[rune(chars[1])]

			if first < second {
				sum += second - first
				i++
			} else {
				sum += first
			}
		}
	}

	return sum
}
