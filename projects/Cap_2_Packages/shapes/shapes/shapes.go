package shapes

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width float64
	Long  float64
}

type Triangle struct {
	Base float64
	Higt float64
}

func (c *Rectangle) Area() float64 {

	return c.Width * c.Long
}
func (c *Triangle) Area() float64 {

	return (c.Base * c.Higt) / 2
}

func (c *Circle) Area() float64 {
	const pi = math.Phi
	return c.Radius * pi * pi
}
