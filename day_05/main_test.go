package main

import "testing"

func TestMain(t *testing.T) {
	fileName := "input_test.txt"
	want := 35
	got := FindLocation(fileName)
	if got != want {
		t.Fatalf("calculated location is incorrect, got %d, want %d", got, want)
	}
}
