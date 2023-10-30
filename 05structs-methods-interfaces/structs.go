package structsmethodsinterfaces

import "math"

type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return t.base * t.height / 2
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

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Length + rectangle.Width) * 2
}

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Length * rectangle.Width
// }
