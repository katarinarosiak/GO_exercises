/*https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

T: O(N)
S: O(1)

- initialize two pointers (buy and sell)
- initialize maxProfit with value of 0
- place buy on the idx 0 and sell on idx 1
- iterate until sell reaches end of the array
- on each iteration:
	- if buy value < sell value:
		- calculate the profit
		if the current profit is higher than maxProfit
			- update maxProfit
		else
			move sell to pointer to the same place as buy
		increment buy
	increment sell (on every iteration)
- return maxProfit

*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(arr)
	fmt.Println(result)
}

func maxProfit(prices []int) int {
	maxProfit := 0
	buy, sell := 0, 1

	for sell < len(prices) {
		if prices[buy] < prices[sell] {
			profit := prices[sell] - prices[buy]
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			buy = sell
		}
		sell++
	}
	return maxProfit
}
