package main

import "fmt"

func inc(i *int) {
	*i++
	fmt.Println(*i)
}

func dec(i *int) {
	*i--
	fmt.Println(*i)
}

func main() {
	i := 0

	go inc(&i)
	go dec(&i)

	i++
	fmt.Println(i)
}
