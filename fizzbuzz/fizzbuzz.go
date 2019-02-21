package main

import (
	"strconv"
)

const (
	Fizz = "Fizz"
	Buzz = "Buzz"
)

func fizzbuzz(i int) string {
	result := ""

	divThree := i%3 == 0
	divFive := i%5 == 0

	if divThree  {
		result = result + Fizz
	}

	if divFive {
		result = result + Buzz
	}

	if !divThree && !divFive {
		result = strconv.Itoa(i)
	}

	return result
}
