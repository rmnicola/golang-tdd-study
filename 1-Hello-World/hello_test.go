package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Hello with a name argument", func(t *testing.T) {
		have := Hello("Chris")
		want := "Hello, Chris!"
    AssertCorrectMsg(t, have, want)
	})

	t.Run("Hello with empty string as a name", func(t *testing.T) {
		have := Hello("")
		want := "Hello, WorlD!"
    AssertCorrectMsg(t, have, want)
	})
}

func AssertCorrectMsg(t testing.TB, have, want string) {
    t.Helper()
		if have != want {
			t.Errorf("\n\thave %q\n\twant %q", have, want)
		}
}
