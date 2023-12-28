package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Hello with a name argument", func(t *testing.T) {
		have := Hello("Chris", "English")
		want := "Hello, Chris!"
		assertCorrectMsg(t, have, want)
	})

	t.Run("Hello with empty string as a name", func(t *testing.T) {
		have := Hello("", "English")
		want := "Hello, World!"
		assertCorrectMsg(t, have, want)
	})

	t.Run("Hello with empty string as a lang", func(t *testing.T) {
		have := Hello("John", "")
		want := "Hello, John!"
		assertCorrectMsg(t, have, want)
	})

	t.Run("Hello in portuguese", func(t *testing.T) {
		have := Hello("João", "Portuguese")
		want := "Olá, João!"
		assertCorrectMsg(t, have, want)
	})

	t.Run("Hello in french", func(t *testing.T) {
		have := Hello("Jacques", "French")
		want := "Bonjour, Jacques!"
		assertCorrectMsg(t, have, want)
	})

}

func assertCorrectMsg(t testing.TB, have, want string) {
	t.Helper()
	if have != want {
		t.Errorf("\n\thave %q\n\twant %q", have, want)
	}
}
