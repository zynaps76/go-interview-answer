package main

import (
	"strconv"
)

const (
	Fizz = "Fizz"
	Buzz = "Buzz"
)

func fizzbuzz(i int) string {
	divThree := i%3 == 0
	divFive := i%5 == 0

	if divThree && divFive {
		return Fizz + Buzz
	} else if divThree {
		return Fizz
	} else if divFive {
		return Buzz
	}

	return strconv.Itoa(i)
}
