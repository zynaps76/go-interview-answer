package main

import (
	"strconv"
)

const (
	Fizz = "Fizz"
	Buzz = "Buzz"
	FizzBuzz = Fizz + Buzz
)

func fizzbuzz(i int) string {
	result := ""

	divFree := i%3 == 0
	divFive := i%5 == 0

	if divFree  {
		result = result + Fizz
	}

	if divFive {
		result = result + Buzz
	}

	if !divFree && !divFive {
		result = strconv.Itoa(i)
	}

	return result
}
