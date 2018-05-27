package main

import "fmt"

func updataSlice(s []int){
	s[0]=100
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println("arr[2:6]=",arr[2:6])
	fmt.Println("arr[:6]=",arr[:6])
	s1 :=arr[2:]
	fmt.Println("s1=",s1)
	s2 := arr[:]
	fmt.Println("s2=",s2)


	fmt.Println("After updateSlice(s1)")
	updataSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updataSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)


	fmt.Println("Reslice")
	fmt.Println(s2)
	s2=s2[:5]
	fmt.Println(s2)
	s2=s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr1 := [...] int {0,1,2,3,4,5,6}
	s3 := arr1[2:6]
	s4 :=s3[3:5]
	fmt.Printf("s3 = %v , len(s3) = %d , cap(s3) = %d \n",s3,len(s3),cap(s3))
	fmt.Printf("s4 = %v , len(s4) = %d , cap(s4) = %d \n",s4,len(s4),cap(s4))
	fmt.Println(s3)
	fmt.Println(s4)
}
