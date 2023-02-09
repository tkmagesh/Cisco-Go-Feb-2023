package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Are() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Length float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Length + r.Width)
}

/*
func PrintArea(x interface{ Area() float32 }) {
	fmt.Println("Area :", x.Area())
}
*/

type AreaCalculator interface {
	Area() float32
}

func PrintArea(x AreaCalculator) {
	fmt.Println("Area :", x.Area())
}

type PerimeterCalculator interface {
	Perimeter() float32
}

func PrintPerimeter(x PerimeterCalculator) {
	fmt.Println("Perimeter :", x.Perimeter())
}

/*
func PrintShape(x interface {
	Area() float32
	Perimeter() float32
}) {
	PrintArea(x)
	PrintPerimeter(x)
}
*/

/*
func PrintShape(x interface {
	AreaCalculator
	PerimeterCalculator
}) {
	PrintArea(x)
	PrintPerimeter(x)
}
*/

//interface composition
type ShapeStatsCalculator interface {
	AreaCalculator
	PerimeterCalculator
}

func PrintShape(x ShapeStatsCalculator) {
	PrintArea(x)
	PrintPerimeter(x)
}

func main() {
	c := Circle{12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{10, 23}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
}
