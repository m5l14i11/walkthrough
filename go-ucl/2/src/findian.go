package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Enter string: ")
	fmt.Scan(&input)

	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.ContainsAny(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
