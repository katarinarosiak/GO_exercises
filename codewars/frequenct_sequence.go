package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	value := FreqSeq("hello world", "-") //"1-1-3-3-2-1-1-2-1-3-1"
	fmt.Println(value)
}

/*



 */

func FreqSeq(str string, sep string) string {
	var final []string

	for _, char := range str {
		num := strconv.Itoa(strings.Count(str, string(char)))
		final = append(final, num)
	}

	return strings.Join(final, sep)
}
