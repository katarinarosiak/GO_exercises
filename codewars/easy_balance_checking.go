package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// checkBook:="1000.00!=\n125 Market !=:125.45\n126 Hardware =34.95\n
	// 127 Video! 7.45\n
	// 128 Book :14.32\n
	// 129 Gasoline ::16.10\n
	// "
	checkBook := "1000.00!=\n\n125 Market !=:125.45\n126 Hardware =34.95\n127 Video! 7.45\n128 Book :14.32\n129 Gasoline ::16.10"
	value := Balance(checkBook) //{4,3,1}
	fmt.Println(value)
}

/*
i: string
o: string


- leave only letters, numbers, dots and spaces
- add "Original Balance:" to the first line "
-  Total_expense__198.27
Average_expense__39.65"
- pn each ;ine add new balance

- some line may be blank
- round to 2 decimals
- no trailing 0
- line separation \n



A:
separate lines into a array of stings
count balance:
	var balance
	val totalExpenses
	- grab the original balance
	- itearet through the arr starting from indx 1
	- from teh string get the price
	- balance - price and save in variable = balance
	- add to variable totalExpenses
	- add to that string at the end " Balance ${balance}"
to the first strings add: "Original Balnace: "
push "Total expense  ${totalExpense}"
push "Average_expense  ${countAverage}"
join with \n new linie char
return
*/

func Balance(book string) string {
	r, _ := regexp.Compile("[^a-zA-Z0-9\n\\.\\ ]")
	book = r.ReplaceAllString(book, "")
	strArr := strings.Split(book, "\n")
	strArr = Clean(strArr)

	originalBalance, _ := strconv.ParseFloat(strArr[0], 64)
	balance := originalBalance
	var totalExpense float64

	for idx := 1; idx < len(strArr); idx++ {
		rgx, _ := regexp.Compile("[0-9]+\\.[0-9]+")
		price, _ := strconv.ParseFloat(rgx.FindString(strArr[idx]), 64)
		balance -= price
		totalExpense += price
		strArr[idx] = strArr[idx] + " Balance " + fmt.Sprintf("%.2f", balance)
	}
	avg := totalExpense / float64(len(strArr)-1)
	strArr[0] = "Original Balance: " + fmt.Sprintf("%.2f", originalBalance)
	strArr = append(strArr, "Total expense  "+fmt.Sprintf("%.2f", totalExpense))
	strArr = append(strArr, "Average expense  "+fmt.Sprintf("%.2f", avg))
	return strings.Join(strArr, "\n")
}

func Clean(arr []string) []string {
	final := []string{}
	for _, i := range arr {
		if i != "" {
			final = append(final, i)
		}
	}
	return final
}
