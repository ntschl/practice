package main

import "testing"

// file name must be xxx_test.go
// test function name starts with Test
// test function takes one arg: t *testing.T
// must import "testing"

// TEST ONE
// func TestHello(t *testing.T) {
// 	got := Hello()
// 	want := "Hello, world!"

// 	if got != want {
// 		t.Errorf("Got %q, want %q", got, want)
// 	}
// }

// TEST TWO
// func TestHello(t *testing.T) {
// 	got := Hello("Chris")
// 	want := "Hello, Chris"

// 	if got != want {
// 		t.Errorf("Got %q, want %q", got, want)
// 	}
// }

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Nate")
		want := "Hello, Nate"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying 'hello world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
