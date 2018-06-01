package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

var a, b, c = 5, 6, "short"

func main() {
	fmt.Println(a, b, c)
	euler()
	triangle()
	consts()
	enums()
}
func euler() {
	fmt.Printf(" %.3f \n",
		cmplx.Exp(1i*math.Pi)+1)
}
func triangle() {
	var a, b int = 3, 4
	fmt.Println(CalcTriangle(a, b))
}

func CalcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const filenmae = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filenmae, c)
}
func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		tb
		pb
	)
	fmt.Println(b, kb, mb, tb, pb)

	fmt.Println(cpp, java, python, golang)

}
