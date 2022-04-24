package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	value := test("xyaabbbccccdefww", "xxxxyyyyabklmopq") //edcba
	fmt.Println(value)
	fmt.Println("asdad" + "eded")
}

/*

i: 2 str
o: new! sorted string
- only letters a-z

- returned strin has to be longest possible
- with no duplicates
- coming from s1 or s2

arr
iterate through first str
take a letter and append to arr if arr doesn't have it already
do teh same with teh other arr
(coerce to string)

sort the array
join and return


*/

func test(s1, s2 string) string {
	arr := []string{}
	strArr := strings.Split(s1+s2, "")
	sort.Strings(strArr)

	for _, lett := range strArr {
		if len(arr) == 0 || arr[len(arr)-1] != lett {
			arr = append(arr, lett)
		}
	}
	return strings.Join(arr, "")
}
