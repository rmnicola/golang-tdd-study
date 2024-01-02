package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		have := Sum(numbers)
		want := 6
		if have != want {
			t.Errorf("Have %d, want %d. Numbers: %v", have, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("calculate sum of two slices", func(t *testing.T) {
		sliceA := []int{1, 2}
		sliceB := []int{0, 9}
		have := SumAll(sliceA, sliceB)
		want := []int{3, 9}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("Have %d, want %d. A: %v, B: %v", have, want, sliceA, sliceB)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, have, want []int) {
		if !reflect.DeepEqual(have, want) {
			t.Errorf("Have %v, want %v", have, want)
		}
	}

	t.Run("calculate sum of two slice tails", func(t *testing.T) {
		sliceA := []int{1, 2, 3, 4, 5}
		sliceB := []int{0, 9, 1, 2, 3}
		have := SumAllTails(sliceA, sliceB)
		want := []int{14, 15}
		checkSums(t, have, want)
	})

	t.Run("calculate sum of two tails with empty slice", func(t *testing.T) {
		sliceA := []int{}
		sliceB := []int{0, 9, 1, 2, 3}
		have := SumAllTails(sliceA, sliceB)
		want := []int{0, 15}
		checkSums(t, have, want)
	})
}
