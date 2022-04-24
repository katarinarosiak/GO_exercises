package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// arr := []int{5, -10, -6, 9, 4}
	value := PlayPass("7U#Y Z9#1 F8TLGQ GG JK77I#9AGWC QCYFB8OM1G", 5) //2008
	fmt.Println(value)
}

/*
	i: string and int
	o: string

	r:
	- shift each letter by given number (if 1 => A=>B)
	- case matters
	- replace each digit by ot's compellne to 9 => (9-num)
	- spaces, punctuation stays the same
	- downcase each letter in odd position
	- upcase letters on even positions (start from 9_)
	- reverse the result

	- iterate through characters
*/

func PlayPass(s string, n int) string {
	final := ""
	for idx, char := range s {
		if unicode.IsLetter(char) {
			letter := string(int(char))
			letter = rotate(letter, n)
			if idx%2 == 0 {
				final += strings.ToUpper(letter)
			} else {
				final += strings.ToLower(letter)
			}

		} else if unicode.IsNumber(char) {
			num, _ := strconv.Atoi(string(char))
			final += strconv.Itoa(9 - num)
		} else {
			final += string(char)
		}
	}

	return reverse(final)
}

func rotate(s string, i int) string {
	alpha := "abcdefghijklmnoprstuvwxyz"
	alphaLen := len(alpha) - 1
	idx := 0

	if strings.ToUpper(s) == s {
		upper := strings.ToUpper(alpha)
		idx = strings.Index(upper, s)
		if idx+i > alphaLen {
			return string(upper[alphaLen-idx])
		}
	} else {
		idx = strings.Index(alpha, s)
		if idx+i > alphaLen {
			return string(alpha[alphaLen-idx])
		}
	}

	return string(alpha[idx+i])
}

func reverse(s string) string {
	reversed := ""
	for idx := len(s) - 1; idx >= 0; idx-- {
		reversed += string(s[idx])
	}
	return reversed
}
