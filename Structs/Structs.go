package structs

import (
	"math"
)

// NOTE : perimeter and area functions could be used to calculate for shapes uncomaptible leading to
// results . Create a struct rectangle

type Rectangle struct {
	width  float64
	height float64
}
type Circle struct {
	radius float64
}
type Triangle struct {
	Height float64
	Base   float64
}

// func Perimeter(width float64, height float64) float64 {
// 	return 2 * (width + height)
// }

// func Area(width float64, height float64) float64 {
// 	return height * width
// }

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.height + rectangle.width)
}

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.height * rectangle.width
// }

// Use of method : https://go.dev/ref/spec#Method_declarations
// A method is a function with a receiver. A method declaration binds an identifier,
//the method name, to a method, and associates the method with the receiver's base type.

//t.Errorf we are calling the method Errorf on the instance of our t (testing.T).

// func (receiverName ReceiverType) MethodName(args).

func (r Rectangle) Area() float64 {
	// It is a convention in Go to have the receiver variable be the first letter of the type.
	// r Rectangle
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (t Triangle) Area() float64 {
	return (t.Base)*(t.Height) * 0.5
}

// Interfaces :	https://golang.org/ref/spec#Interface_types
//allow you to make functions that can be used with different types and create highly-decoupled code
// whilst still maintaining type-safety.

//Go interface resolution is implicit. If the type you pass in matches
//what the interface is asking for, it will compile.
// so Method for rectangle and circle have an area func with flaot64 thus it passes but other not

/*

	There is some duplication in our tests.
	All we want to do is take a collection of shapes, call the Area() method on them and then check the result.
	some kind of checkArea function that we can pass both Rectangles and Circles to,
	but fail to compile if we try to pass in something that isn't a shape.


*/
