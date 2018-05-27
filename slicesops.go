package main

import (
	"fmt"
)

func printSlices(s []int)  {
	fmt.Printf("%v,len = %d , cap = %d \n",
		s,len(s),cap(s))
}
func main() {
	var  s [] int  //Zero value for slice is nil

	for i := 0; i < 100 ; i++ {
		printSlices(s)
		s = append(s, 2*i+1)

	}
	fmt.Println(s)

	s1 := [] int {2,4,6,8}
	printSlices(s1)

	s2 := make([]int , 16)

	s3 := make([]int ,10 ,32)

	printSlices(s2)
	printSlices(s3)
}

