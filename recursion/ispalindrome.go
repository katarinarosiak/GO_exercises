package main

import (
	"fmt"
)

func main() {

	value := IsPalindrome("bbbbbc") //true
	fmt.Println(value)
}

/*

- if odd letter i mellan counts as two
- aba => true
- abba => true
- case insensitive


- abcde


- get middle slice
- check if palindrome
- if false return false immidiatelly
- else expand slice and check
*/

func IsPalindrome(str string) bool {

	if len(str) == 2 {
		return str[0] == str[len(str)-1]
	}

	start := len(str)/2 - 1
	end := len(str)/2 + 1
	slice := str[start:end]

	return IsPalindrome(slice) && str[0] == str[len(str)-1]

}
