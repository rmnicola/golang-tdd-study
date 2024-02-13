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

	t.Run("Perimeter Rect", func(t *testing.T) {
		rect := Rectangle{2.0, 3.0}
		checkPerimeter(t, rect, 10.0)
	})

	t.Run("Perimeter Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkPerimeter(t, circle, 62.83)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		have := shape.Area()
		assertCorrectFloat(t, have, want)
	}

	t.Run("Area Rect", func(t *testing.T) {
		rect := Rectangle{2.0, 3.0}
		checkArea(t, rect, 6.0)
	})

	t.Run("Area Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.15)
	})
}

func assertCorrectFloat(t *testing.T, have, want float64) {
	t.Helper()
	absError := math.Abs(have - want)
	if absError > equalityThreshold {
		t.Errorf("Have %g Want %g. Error %g\n", have, want, absError)
	}
}
