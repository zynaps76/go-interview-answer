package main

import (
	"strconv"
)

const (
	Fizz = "Fizz"
	Buzz = "Buzz"
)

func fizzbuzz(i int) string {
	fizz := i%3 == 0
	buzz := i%5 == 0

	if fizz && buzz {
		return Fizz + Buzz
	} else if fizz {
		return Fizz
	} else if buzz {
		return Buzz
	}

	return strconv.Itoa(i)
}
