// package main

// import (
// 	"fmt"
// 	// "strings"
// )

// // func validateBanner(input string) map[rune]string {

// // }

// func main() {
// 	d := map[int] string{
// 		4 : "food",
// 		5 : "spoon",
// 		6 : "water",
// 	}
// 	// delete(d, 4)
// 	// fmt.Println(d)
// 	value, ok := d[4]
// 	if ok {
// 		fmt.Println("true")
// 	}else if !ok {
// 		fmt.Println(value)
// 	}

// }

package main

import "fmt"

func validateBanner(input string) map[rune]int {
	LIMIT := 3
	charCounts := make(map[rune]int)
	good := 43
	current := rune(36)
	// count := 0
	for _, ch := range input {
		if ch == current {
			charCounts[ch] += good
			// return charCounts
		}
	}
	charCounts[current] = LIMIT
	current++
		
		

	// count++

	// charrune := make(map[rune]string)

	return charCounts
}
func main() {
	fmt.Println(validateBanner(""))
}
