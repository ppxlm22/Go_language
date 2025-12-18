package main

import (
	"math"
	// "fmt"
)
type shape interface{
	Area() float64
}
type Rectangle struct{
	leaght,width float64
}
type Triangle struct{
	base,heigh float64
}
type Circle struct{
	radius float64
}
func (r Rectangle) Area() float64{
	return r.leaght * r.width
}

func (t Triangle) Area() float64{
	return 0.5* t.base * t.heigh
}
func (c Circle) Area() float64{
	return math.Pi * c.radius *c.radius
}
func calculateArea(s shape) float64{
	return s.Area()
}