package shapes

import (
	"math"
	"testing"
)

const equalityThreshold = 0.01

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		have := shape.Perimeter()
		assertCorrectFloat(t, have, want)
	}

	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{2.0, 3.0}, 10.0},
		{Circle{10.0}, 62.83},
	}

	for _, testCase := range perimeterTests {
		checkPerimeter(t, testCase.shape, testCase.want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		have := shape.Area()
		assertCorrectFloat(t, have, want)
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{2.0, 3.0}, 6.0},
		{Circle{10.0}, 314.15},
		{Triangle{2.0, 3.0}, 3.0},
	}

	for _, testCase := range areaTests {
		checkArea(t, testCase.shape, testCase.want)
	}
}

func assertCorrectFloat(t *testing.T, have, want float64) {
	t.Helper()
	absError := math.Abs(have - want)
	if absError > equalityThreshold {
		t.Errorf("Have %g Want %g. Error %g\n", have, want, absError)
	}
}
