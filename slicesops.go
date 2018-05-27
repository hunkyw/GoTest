package main

import "fmt"

func printSlices(s []int)  {
	fmt.Printf("len = %d , cap = %d \n",len(s),cap(s))
}
func main() {
	var  s [] int  //Zero value for slice is nil

	for i := 0; i < 100 ; i++ {
		printSlices(s)
		s = append(s, 2*i+1)

	}
	fmt.Println(s)
}

