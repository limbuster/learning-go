package main

import "math"

// Shape defines the behavior of different shapes
type Shape interface {
	Area() float64
}

// Rectangle defines rectangular shape with width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of the rectangle
func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns the rectangle
func (rectangle Rectangle) Area() float64 {
	return rectangle.Height * rectangle.Width
}

// Circle defines circular shape with radius
type Circle struct {
	Radius float64
}

// Perimeter returns the circumference
func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

// Area returns the area of the circle
func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

// Triangle defines triangular shape with base and height
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter returns the perimeter of the triangle
func (triangle Triangle) Perimeter() float64 {
	return 0
}

// Area returns the area of the triangle
func (triangle Triangle) Area() float64 {
	return 0.5 * triangle.Height * triangle.Base
}
