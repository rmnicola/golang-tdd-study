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
		{shape: Rectangle{Width: 2.0, Height: 3.0}, want: 10.0},
		{shape: Circle{Radius: 10.0}, want: 62.83},
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
		{shape: Rectangle{Width: 2.0, Height: 3.0}, want: 6.0},
		{shape: Circle{Radius: 10.0}, want: 314.15},
		{shape: Triangle{Base: 2.0, Height: 3.0}, want: 3.0},
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
