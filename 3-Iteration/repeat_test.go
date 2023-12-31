package iteration

import "testing"

func TestRepeat(t *testing.T) {
	have := Repeat("a")
	want := "aaaaa"
  assertCorrectOutput(t, have, want)
}

func assertCorrectOutput(t testing.TB, have, want string) {
	t.Helper()
	if have != want {
		t.Errorf("\n\thave %q\n\twant %q", have, want)
	}
}
