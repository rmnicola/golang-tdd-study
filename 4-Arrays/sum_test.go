package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		have := Sum(numbers)
		want := 6
		assertCorrectSum(t, have, want, numbers)
	})

}

func assertCorrectSum(t testing.TB, have, want int, numbers []int) {
	t.Helper()
	if have != want {
		t.Errorf("Have %d, want %d. Numbers: %v", have, want, numbers)
	}
}
