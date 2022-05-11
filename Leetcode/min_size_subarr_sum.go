package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	value := minSubArrayLen(11, a) //2008
	fmt.Println(value)
}

/*NOT DONE YET

https://leetcode.com/problems/minimum-size-subarray-sum/


i: arr of +int
int + (target)

r: int (min len of contiguous subarr)
 => of which sum >= target
 or 0 if not found

r:
contiguous subarr => adjacent


w=1
size = 3
sum  7

2,3,1,2,4,3  7
      a
          r





*/

func minSubArrayLen(target int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	left := nums[0]
	right := left + 1

	return 0
}

func sum(arr []int) int {
	result := 0
	for _, num := range arr {
		result += num
	}
	return result
}
