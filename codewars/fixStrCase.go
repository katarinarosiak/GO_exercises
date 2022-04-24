package main

import (
	"fmt"
	"strings"
)

func main() {

	value := solve("CODe") //"CODE"
	fmt.Println(value)
}

/*
	if more upper change all to upper
	if more lower change all to lower

	- itertae through string
	count number of lower and upper


*/

func solve(str string) string {
	lower := 0
	upper := 0

	for i := 0; i < len(str); i++ {
		if string(str[i]) == strings.ToUpper(string(str[i])) {
			upper++
		} else {
			lower++
		}
	}

	if upper > lower {
		return strings.ToUpper(str)
	} else {
		return strings.ToLower(str)
	}

}
