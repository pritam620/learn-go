package structs

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.length)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.length * rectangle.width
}

type Rectangle struct {
	width, length float64
}
