package main

import "fmt"

func lengthOfNonrepeatingSubStr(s string) int{
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s)  {
		lastI,ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastOccurred[ch]+1

		}
		if i - start + 1 > maxLength{

			maxLength = i -start +1
		}
		lastOccurred[ch]=i
	}
	return maxLength
}
func main() {
	fmt.Println(lengthOfNonrepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonrepeatingSubStr("bbbbbbb"))
	fmt.Println(lengthOfNonrepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonrepeatingSubStr(""))
	fmt.Println(lengthOfNonrepeatingSubStr("b"))
	fmt.Println(lengthOfNonrepeatingSubStr("abcdefghijk"))
}
