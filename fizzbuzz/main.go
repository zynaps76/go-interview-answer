package main

/*
Напишите программу, которая выводит на экран числа от 1 до 100.
При этом вместо чисел, кратных трем, программа должна выводить слово «Fizz»,
аx вместо чисел, кратных пяти — слово «Buzz». Если число кратно и 3, и 5,
то программа должна выводить слово «FizzBuzz»
*/

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
