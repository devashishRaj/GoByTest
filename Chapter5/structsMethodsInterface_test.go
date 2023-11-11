package structsMethodsInterface

import "testing"

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.00
	if got != want {
		t.Errorf("Got %.2f want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		/*
// 			helper does not need to concern itself with whether the shape is a Rectangle or a Circle or
// 			a Triangle. By declaring an interface, the helper is decoupled from the concrete types and
// 			only has the method it needs to do its job.
// 		*/

// 		// NOTE :  using interfaces to declare only what you need is very important in software design
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("Got %.2f wanted %.2f", got, want)
// 		}
// 	}
// 	t.Run("Calculate rectangle area", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		checkArea(t, rectangle, 72.0)
// 	})
// 	t.Run("Calculate circle area", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

/*

	In terms of significant digits:
	Accuracy doesn't necessarily depend on the number of significant digits.
	It is about how close the measured value is to the true value.

	Precision often correlates with the number of significant digits.
	More precise measurements tend to have more significant digits because they provide more
	information about the variation or uncertainty in the measurements.

*/

// Table driven tests : https://github.com/golang/go/wiki/TableDrivenTests

func TestArea(t *testing.T) {
	/*
		"anonymous struct", areaTests. We are declaring a slice of structs by using []struct with
		two fields, the shape and the want. Then we fill the slice with cases.
		 They are a great fit when you wish to test various implementations of an interface, or
		if the data being passed in to a function has lots of different requirements that need testing.
	*/
	//It's not clear what all the numbers represent and  tests should be easily understandable.

	// areaTests := []struct {
	// 	shape Shape
	// 	want  float64
	// }{
	// 	{Rectangle{12.0, 6.0}, 72.0},
	// 	{Circle{10.0}, 314.1592653589793},
	// 	// this test will fail , but look how easy it is add a new test
	// 	//{Rectangle{12.0 , 6.0 }, 36.0},
	// 	// Adding new test
	// 	{Triangle{12.0 , 6.0 } , 36.0},
	// }

	// you can optionally name the fields.

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{height: 6.0, width: 12.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12.0, Height: 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		//By wrapping each case in a t.Run you will have clearer test output on failures as it will 
		//print the name of the case
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				// We can change our error message into %#v got %g want %g.
				//The %#v format string will print out our struct with the values in its field,
				//so the developer can see at a glance the properties that are being tested.
				t.Errorf("%#v got %g want %g" , tt.shape , got , tt.hasArea)
			}
		})
	}
}
