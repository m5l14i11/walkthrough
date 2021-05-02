package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	arr := make([]int, 0, 3)

	for {
		var input string

		fmt.Scanln(&input)

		if input == "X" {
			break
		}

		n, err := strconv.Atoi(input)

		if err != nil {
			break
		}

		arr = append(arr, n)

		sort.Ints(arr)

		fmt.Printf("%d \n", arr)
	}
}
