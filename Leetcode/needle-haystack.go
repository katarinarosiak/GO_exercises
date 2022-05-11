//https://leetcode.com/problems/implement-strstr/submissions/

package main

import (
	"fmt"
)

func main() {
	haystack := "asadvmmv"
	needle := ""
	value := strStr(haystack, needle)
	fmt.Println(value)
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	needleLen := len(needle)
	if needleLen <= len(haystack) {
		for i := 0; i <= len(haystack)-needleLen; i++ {
			if haystack[i:i+needleLen] == needle {
				return i
			}
		}
	}

	return -1
}
