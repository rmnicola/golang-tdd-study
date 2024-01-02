package sum

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	have := Sum(numbers)
	want := 15

	if have != want {
		t.Errorf("Have %d, want %d. Numbers: %v", have, want, numbers)
	}
}
