package shapes

import "math"

// Shape is a shape
type Shape interface {
	Area() float64
}

// Rectangle describes the shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle describes the shape
type Circle struct {
	Radius float64
}

// Triangle describes the shape
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter gives the perimeter of the shape
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area gives the area of the shape
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}

// Area gives the area of the shape
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Area gives the area of the shape
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area gives the area of the shape
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
