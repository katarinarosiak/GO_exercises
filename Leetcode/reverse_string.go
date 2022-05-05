/*
https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/879/

i: array of characters (bytes)
o: no output=> reverse array of bytes (reversed)

- in place
- O(1) extra memory

- len at least one char , can be large
s[i] is printable ASCII char


iterate through array X len arr
y => len arr -1
- swap two elements arr[i] <=> arr[y]
- y--, i++


[1, 2, 3, 4, 5]
*/

package main

func main() {
	arr := []byte{'h', 'a', 'e', 'l'}

	reverseString(arr)

}

func reverseString(s []byte) {
	for i, y := 0, len(s)-1; i < len(s)/2; i++ {
		s[y], s[i] = s[i], s[y]
		y--
	}
}
