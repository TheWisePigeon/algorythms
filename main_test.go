package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestBinarySearch(t *testing.T) {
	containsValue := []int{1, 2, 3, 4, 5, 6, 7, 8}
	doesntContainValue := []int{1, 2, 3, 5, 7, 8, 9, 10, 12}

	ok, index := BinarySearch(containsValue, 2)
	if !ok {
		t.Fatalf("First test failed")
	}
	fmt.Println("Found needle at index:", index)
	ok, _ = BinarySearch(doesntContainValue, 4)
	if ok {
		t.Fatalf("Second test failed")
	}
}
