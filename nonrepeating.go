package main

import "fmt"

func lengthOfNonrepeatingSubStr(s string) int{
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s)  {
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
	fmt.Println(lengthOfNonrepeatingSubStr("这里是慕课网"))
	fmt.Println(lengthOfNonrepeatingSubStr("一二三三二一"))
	fmt.Println(lengthOfNonrepeatingSubStr(
		"黑化肥会挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
