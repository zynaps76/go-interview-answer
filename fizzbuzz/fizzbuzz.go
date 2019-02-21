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

	if i%3 == 0 {
		result = result + Fizz
	}

	if i%5 == 0 {
		result = result + Buzz
	}

	if result == "" {
		result = strconv.Itoa(i)
	}

	return result
}
