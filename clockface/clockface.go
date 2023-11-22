package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}
const (

	secondHandLength = 90
	clockCentreX = 150
	clockCentreY = 150
)
// func SecondHand(t time.Time) Point {

// 	/*
// 		1.Scale it to the length of the hand
// 		2.Flip it over the X axis to account for the SVG having an origin in the top left hand corner
// 		3.Translate it to the right position (so that it's coming from an origin of (150,150))
	
// 	*/
// 	p := SecondHandPoint(t)
// 	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
// 	p = Point{p.X, -p.Y}
// 	p = Point{p.X + clockCentreX, p.Y + clockCentreY} //translate
// 	return p
// }

func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func SecondHandPoint(t time.Time) Point{
	angle := SecondsInRadians(t)

	x:= math.Sin(angle)
	y:= math.Cos(angle)
	return Point{x,y}
}
