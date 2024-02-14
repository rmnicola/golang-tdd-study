package shapes

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rect Rectangle) Perimeter() (result float64) {
	result = 2 * (rect.Width + rect.Height)
	return
}

func (rect Rectangle) Area() (result float64) {
	result = rect.Width * rect.Height
	return
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() (result float64) {
	result = math.Pi * 2 * c.Radius
	return
}

func (c Circle) Area() (result float64) {
	result = math.Pi * c.Radius * c.Radius
	return
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() (result float64) {
	result = (t.Base * t.Height) / 2
	return
}

func (t Triangle) Perimeter() (result float64) {
	return
}
