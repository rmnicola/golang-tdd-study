package main

import "testing"

func TestHello(t *testing.T) {
	have := Hello("Chris")
	want := "Hello, Chris"

	if have != want {
		t.Errorf("\n\thave %q\n\twant %q", have, want)
	}
}
