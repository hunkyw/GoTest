package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"io"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)

}
func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

	}
}
func main() {
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //1101
		convertToBin(72387885),
		convertToBin(0))
	printFile("abc.txt")
	s := `abc"d"
	kkk
	123
	
	p`
	printFileContents(strings.NewReader(s))
}
