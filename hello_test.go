package main

import "testing"

// file name must be xxx_test.go
// test function name starts with Test
// test function takes one arg: t *testing.T
// must import "testing"

// func TestHello(t *testing.T) {
// 	got := Hello()
// 	want := "Hello, world!"

// 	if got != want {
// 		t.Errorf("Got %q, want %q", got, want)
// 	}
// }

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
