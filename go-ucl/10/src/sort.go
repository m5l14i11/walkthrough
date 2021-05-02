package main

import (
	"fmt"
	"sort"
	"sync"
)

func Sort(wg *sync.WaitGroup, arr []int) {
	sort.Ints(arr)

	fmt.Println(arr)

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	arr := make([]int)

	go Sort(&wg, arr)
}
