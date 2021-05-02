package main

import "fmt"

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var input int

	arr := make([]int, 0, 10)

	fmt.Println("Prompt 10 integers: ")

	for {
		fmt.Scan(&input)
		arr = append(arr, input)

		if len(arr) >= 10 {
			break
		}
	}

	BubbleSort(arr)

	fmt.Printf("%d \n", arr)
}
