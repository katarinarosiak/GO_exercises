package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	var a1 = []int{2, 2, 3}
	var a2 = []int{4, 9, 9}

	value := Comp(a1, a2)
	fmt.Println(value)

}

/*
i: arra a arr b
o: bool

- check if two arrays have the same elemes
and the same muktiplicies


- multiplicy:
number of times it appears

- same:
- el in b are the elements in a squared

- order doesn't matter

- might be array or {}
- if null or nothing return false
- have the same size

a = [121, 144, 19, 161, 19, 144, 19, 11]
b = [121, 14641, 20736, 361, 25921, 361, 20736, 361]

iteare through b,
take el square and map into {}
b*b : true

iterate through a
if any of the number doesn't exist in a return false immidietely
else return true
*/

func Comp(array1 []int, array2 []int) bool {
	a := []int{}

	if len(array1) != len(array2) || array1 == nil || array2 == nil {
		return false
	}

	for _, num := range array1 {
		a = append(a, num*num)
	}

	sort.Ints(a)
	sort.Ints(array2)

	return reflect.DeepEqual(a, array2)
}
