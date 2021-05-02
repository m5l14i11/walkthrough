package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var input string

	user := make(map[string]string)

	fmt.Scanln(&input)
	user["name"] = input

	fmt.Scanln(&input)
	user["address"] = input

	data, _ := json.Marshal(user)

	fmt.Printf("%s \n", string(data))
}
