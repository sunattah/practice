package main

import (
	"fmt"
	"strconv"
)

func Atoi(s string) int {
	file, err := strconv.Atoi(s)
	if err != nil {
		return file
	}
	return file

}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
