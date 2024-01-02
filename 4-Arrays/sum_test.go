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
