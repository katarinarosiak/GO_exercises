package main

import (
	"fmt"
)

func main() {
	var a1 = []int{2, 4, 0, 100, 4, 11, 2602, 36}
	// var a2 = []string{"bcc", "bccc", "lade", "lade", "abada"}

	value := FindOutlier(a1)
	fmt.Println(value)

}
