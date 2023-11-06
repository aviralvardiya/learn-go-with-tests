package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Raman")

	got := buffer.String()
	want := "Hello, Raman"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}