package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

func main() {
	var acceleration float64
	var velocity float64
	var displacement float64
	var time float64

	fmt.Println("Enter acceleration: ")
	fmt.Scanln(&acceleration)

	fmt.Println("Enter velocity: ")
	fmt.Scanln(&velocity)

	fmt.Println("Enter displacement: ")
	fmt.Scanln(&displacement)

	fn := GenDisplaceFn(acceleration, velocity, displacement)

	fmt.Println("Enter time: ")
	fmt.Scanln(&time)

	fmt.Println(fn(time))
}
