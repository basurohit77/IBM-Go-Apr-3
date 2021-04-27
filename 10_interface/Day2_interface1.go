package main

import (
	"fmt"
	"math"
)

type Dimension interface {
	IHaveArea
	IHavePerimeter
}

// Area is common to both struct, so the beheavior of both is same, hence we creat an interface with area function.
// Interface can be used by both the struct
type IHaveArea interface {
	Area() float32
}

type IHavePerimeter interface {
	Perimeter() float32
}

type Rect struct {
	length, width float32
}

type Circle struct {
	radius float32
}


func (r Rect) Area() float32 {
	return r.length * r.width
}

func (r Rect) Perimeter() float32 {
	return 2*r.length + 2*r.width
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func main() {
	rect := Rect{10, 20}
	circle := Circle{15}

	// Normal way of calling
	var ra = rect.Area()
	fmt.Println("area of rect  ", ra)

	var ca = circle.Area()
	fmt.Println("area of circle  ", ca)

	/*
	    // While calling the interface we can pass struct object ==> As the interface is applicable for both the  struct
	    // By this  if more number of struct are their and the behivour is same we creat one interface for all
		printArea(rect)
		printArea(circle)
		printPerimeter(rect)
		printPerimeter(circle)
	*/

	fmt.Println(rect)
	printDimension(rect)
	printDimension(circle)
}

// creat a function and pass parameter as interface  containing Area function
func printArea(o IHaveArea) {
	fmt.Println("Area = ", o.Area())
}

func printPerimeter(o IHavePerimeter) {
	fmt.Println("Perimeter = ", o.Perimeter())
}

func printDimension(d Dimension) {
	printArea(d)
	printPerimeter(d)
}

func (r Rect) String() string {
	return fmt.Sprintf("React {length=%v, width=%v, area=%v, perimeter=%v}", r.length, r.width, r.Area(), r.Perimeter())
}

/*
	perimeter for rect = 2 * length + 2 * width
	perimter for circle = 2 * pi * radius
*/