package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	seq := make([]byte, 16)

	rand.Read(seq)
	fmt.Println("Unsorted:", seq)

	bubbleSort(seq)
	fmt.Println("Sorted:", seq)
}
