package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type rectangle struct {
	length, breath float64
}
type circle struct {
	radious float64
}
type square struct {
	a float64
}

func (r rectangle) area() float64 {
	return r.length * r.breath
}
func (c circle) area() float64 {
	return math.Pi * c.radious * c.radious
}
func (sq square) area() float64 {
	return sq.a * sq.a
}
func calculate(s shape) float64 {
	return s.area()
}
func main() {
	r := rectangle{21, 12}
	c := circle{10}
	sq := square{28}

	rectangle_area := calculate(r)
	circle_area := calculate(c)
	square_area := calculate(sq)

	fmt.Println("area of rectangle = ", rectangle_area)
	fmt.Println("area of circle = ", circle_area)
	fmt.Println("area of square = ", square_area)
}
