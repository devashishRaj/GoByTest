package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// instead of x int , y int we can use x,y int
// not using named return value as use of function is pretty clear here , to add.
// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
	// what if we just return 4 then it's quite a game of cat and mouse eh!
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

// Please note that the example function will not be executed if you remove the comment // Output: 6.
//Although the function will be compiled, it won't be executed.
// as with it it will appear in offical documentation as :
// https://pkg.go.dev/github.com/quii/learn-go-with-tests/integers/v2
