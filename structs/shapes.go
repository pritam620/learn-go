package structs

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Length)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Length * rectangle.Width
}

// You can have functions with the same name declared in different packages
// or We can define methods on our defined types instead.
// func Area(circle Circle) float64 {
// return math.Pi * circle.Radius * circle.Radius
// }

type Rectangle struct {
	Width, Length float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}
