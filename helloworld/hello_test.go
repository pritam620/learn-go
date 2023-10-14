package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Greet with given name", func(t *testing.T) {
		got := Hello("Pritam", "")
		want := "Hello, Pritam"
		assertCorrectMessage(got, want, t)
	})
	t.Run("Greet with Hello world if empty name is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(got, want, t)
	})
	t.Run("Greet in Spanish", func(t *testing.T) {
		got := Hello("Pritam", "Spanish")
		want := "Hola, Pritam"
		assertCorrectMessage(got, want, t)
	})
	t.Run("Greet in French", func(t *testing.T) {
		got := Hello("Pritam", "French")
		want := "Bonjour, Pritam"
		assertCorrectMessage(got, want, t)
	})
}

/**
* For helper functions, it's a good idea to accept a testing.TB
* which is an interface that *testing.T and *testing.B both satisfy, so you can call helper functions from a test, or a benchmark
**/
func assertCorrectMessage(got string, want string, t testing.TB) {
	t.Helper() //t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helpe
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
