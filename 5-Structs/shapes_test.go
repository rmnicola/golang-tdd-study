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
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 2.0, Height: 3.0}, want: 10.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, want: 62.83},
	}

	for _, testCase := range perimeterTests {
		t.Run(testCase.name, func(t *testing.T) {
			checkPerimeter(t, testCase.shape, testCase.want)
		})
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		have := shape.Area()
		assertCorrectFloat(t, have, want)
	}

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 2.0, Height: 3.0}, want: 6.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, want: 314.15},
		{name: "Triangle", shape: Triangle{Base: 2.0, Height: 3.0}, want: 3.0},
	}

	for _, testCase := range areaTests {
		t.Run(testCase.name, func(t *testing.T) {
			checkArea(t, testCase.shape, testCase.want)
		})
	}
}

func assertCorrectFloat(t *testing.T, have, want float64) {
	t.Helper()
	absError := math.Abs(have - want)
	if absError > equalityThreshold {
		t.Errorf("Have %g Want %g. Error %g\n", have, want, absError)
	}
}
