package shapes

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) (area float64) {
	return rectangle.Width * rectangle.Height
}
