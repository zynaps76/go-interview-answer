package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	payload := map[int]string{
		1:  "1",
		3:  Fizz,
		4:  "4",
		5:  Buzz,
		15: FizzBuzz,
	}

	for value, expected := range payload {
		result := fizzbuzz(value)
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		fizzbuzz(i)
	}
}
