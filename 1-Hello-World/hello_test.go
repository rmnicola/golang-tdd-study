package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Hello with a name argument", func(t *testing.T) {
		have := Hello("Chris", "English")
		want := "Hello, Chris!"
		AssertCorrectMsg(t, have, want)
	})

	t.Run("Hello with empty string as a name", func(t *testing.T) {
		have := Hello("", "English")
		want := "Hello, World!"
		AssertCorrectMsg(t, have, want)
	})

	t.Run("Hello with empty string as a lang", func(t *testing.T) {
		have := Hello("John", "")
		want := "Hello, John!"
		AssertCorrectMsg(t, have, want)
	})

	t.Run("Hello in portuguese", func(t *testing.T) {
		have := Hello("João", "Portuguese")
		want := "Olá, João!"
		AssertCorrectMsg(t, have, want)
	})

	t.Run("Hello in french", func(t *testing.T) {
		have := Hello("Jacques", "French")
		want := "Bonjour, Jacques!"
		AssertCorrectMsg(t, have, want)
	})

}

func AssertCorrectMsg(t testing.TB, have, want string) {
	t.Helper()
	if have != want {
		t.Errorf("\n\thave %q\n\twant %q", have, want)
	}
}
