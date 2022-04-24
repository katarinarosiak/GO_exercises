package main

import (
	"fmt"
)

func main() {

	value := firstNoDupl("minimumx") //n
	fmt.Println(value)
}

/*
func =>
i = string
o = return first non-dupl char in str

	A:
	- create a hast table of each char and count occurencies
	- iterate trhough the string again and if the char == 0 return the char (to string)


*/

func firstNoDupl(str string) string {
	hashTable := make(map[byte]int)

	for idx := 0; idx < len(str); idx++ {
		if hashTable[str[idx]] >= 1 {
			hashTable[str[idx]]++
		} else {
			hashTable[str[idx]] = 1
		}
	}

	for idx := 0; idx < len(str); idx++ {
		if hashTable[str[idx]] == 1 {
			return string(str[idx])
		}
	}
	return ""
}
