package main

import (
	"fmt"
)

var basketball_players = []map[string]string{
	{"first_name": "Jill", "last_name": "Huang", "team": "Gators"},
	{"first_name": "Janko", "last_name": "Barton", "team": "Sharks"},
	{"first_name": "Wanda", "last_name": "Vakulskas", "team": "Sharks"},
	{"first_name": "Jill", "last_name": "Moloney", "team": "Gators"},
	{"first_name": "Luuk", "last_name": "Watkins", "team": "Gators"},
}

var football_players = []map[string]string{
	{"first_name": "Hanzla", "last_name": "Radosti", "team": "32ers"},
	{"first_name": "Tina", "last_name": "Watkins", "team": "Barleycorns"},
	{"first_name": "Alex", "last_name": "Patel", "team": "32ers"},
	{"first_name": "Jill", "last_name": "Huang", "team": "Barleycorns"},
	{"first_name": "Wanda", "last_name": "Vakulskas", "team": "Barleycorns"},
}

/*
	create a map of name and last name of first arr
	iterate through second
	- take the name and last name, and map
	- if exist in the other add to a final arr
*/

func main() {

	value := players(basketball_players, football_players) //2008
	fmt.Println(value)
}

/*
	i: two slices [{}, {}]
	return arr of names of players that exist in both arrays
	- O(N+M)

*/

func players(basketball_players, football_players []map[string]string) []string {
	final := []string{}
	hash := make(map[string]bool)

	for _, player := range basketball_players {
		name := player["first_name"] + " " + player["last_name"]
		hash[name] = true
	}

	for _, player := range football_players {
		name := player["first_name"] + " " + player["last_name"]

		if hash[name] {
			final = append(final, name)
		}
	}
	fmt.Println(hash)
	return final
}
