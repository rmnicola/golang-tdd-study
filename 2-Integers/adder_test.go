package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
	have := Add(2, 2)
	want := 4
	assertCorrectInt(t, have, want)
}

func ExampleAdd() {
	sum := Add(3, 4)
	fmt.Println(sum)
	// Output: 7
}

func assertCorrectInt(t testing.TB, have, want int) {
	t.Helper()
	if have != want {
		t.Errorf("\n\thave %d\n\twant %d", have, want)
	}
}
