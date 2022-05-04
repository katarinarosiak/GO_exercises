//https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/564/

/*
i: int arr (prices)
where price[i] is a price of given stock on ith day
o: int  (profit)


R:
- if never profit => 0
price >= 0
arr len >= 0
- multiple number of buy and sale


lowest 1
currProfit = 4
finalProfit = 4


[7,1,5,3,6,4]

bought =
profit =
- iterate through the array
- take three numbers:
	buy [0]
	sell [1]



*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	value := maxProfit(arr)
	fmt.Println(value)
}

func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
