package main

import (
	"fmt"
	"strconv"
)

func hexDecimal(hex string) (int64, error) {
	file, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0, err
	}
	return file, nil
}

func main() {
	fmt.Println(hexDecimal("1E"))
	fmt.Println(hexDecimal("FF"))
}
