package main

import (
	"fmt"
	"strings"
)

func main() {

	value := ToNato("Golf Oscar Foxtrot Oscar Romeo India Tango !")
	fmt.Println(value)
}

/*
i: string
output: string (Nato transalted)

- library NATO
- spaces
- punctuation count
- puntuation separated by a space , every word too
- no trailing whitespaces
- returned words capitalized

A:
- make a struct of the array
- get all the first letters and punct
- map the lettes to words
- join with space


*/

func ToNato(sentence string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "Xray", "Yankee", "Zulu", "!", ",", "."}
	var mapNato = map[string]string{}
	final := []string{}

	for _, word := range nato {
		mapNato[string(word[0])] = word
	}

	for _, char := range sentence {
		word := mapNato[strings.ToUpper(string(char))]
		if word != "" {
			fmt.Println(word)
			final = append(final, word)
		}
	}

	return strings.Join(final, " ")
}
