package main

import (
	"errors"
)

func bubbleSort(slice []byte) {
	for y := range slice {
		i, _ := findMax(slice[y:])

		// swap
		slice[y], slice[i+y] = slice[i+y], slice[y]
	}
}

func findMax(slice []byte) (int, error) {
	if len(slice) == 0 {
		return -1, errors.New("empty slice")
	}

	minVal := slice[0]
	index := 0

	if len(slice) > 1 {
		offset := 1
		for i, v := range slice[offset:] {
			if v < minVal {
				minVal = v
				index = i + offset
			}
		}
	}

	return index, nil
}
