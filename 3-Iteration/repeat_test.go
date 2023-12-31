package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat 'a' 3 times", func(t *testing.T) {
		have := Repeat("a", 3)
		want := "aaa"
		assertCorrectOutput(t, have, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func assertCorrectOutput(t testing.TB, have, want string) {
	t.Helper()
	if have != want {
		t.Errorf("\n\thave %q\n\twant %q", have, want)
	}
}
