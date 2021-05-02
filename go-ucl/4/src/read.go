package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func TruncateName(s string, size int) string {
	if len(s) > size {
		return s[:size]
	}
	return s
}

func main() {
	persons := make([]Person, 0, 3)
	var filePath string
	fmt.Println("Enter file path: ")
	fmt.Scanln(&filePath)

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		l := strings.Split(line, " ")

		p := Person{
			fname: TruncateName(l[0], 20),
			lname: TruncateName(l[1], 20),
		}
		persons = append(persons, p)
	}

	for _, person := range persons {
		fmt.Printf("First Name: %s, Last Name: %s \n", person.fname, person.lname)
	}
}
