package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Pritam")
	want := "Hello Pritam"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
