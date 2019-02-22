package main

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func makeRandomSlice(n int) []byte {
	seq := make([]byte, n)
	rand.Read(seq)

	return seq
}

func TestBubbleSort(t *testing.T) {
	payload := []byte{255, 1, 16, 17, 11, 254, 47}
	expected := []byte{1, 11, 16, 17, 47, 254, 255}

	bubbleSort(payload)

	if !bytes.Equal(payload, expected) {
		t.Errorf("Expected %s, got %s", expected, payload)
	}
}

func BenchmarkBubbleSort64(b *testing.B) {
	seq := makeRandomSlice(64)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		bubbleSort(seq)
	}
}

func BenchmarkBubbleSort128(b *testing.B) {
	seq := makeRandomSlice(128)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		bubbleSort(seq)
	}
}

func BenchmarkBubbleSort256(b *testing.B) {
	seq := makeRandomSlice(256)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		bubbleSort(seq)
	}
}
