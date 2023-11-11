package arraySlices

import (
	"reflect"
	"testing"
)
// https://go.dev/blog/slices-intro

/*
An interesting property of arrays is that the size is encoded in its type. If you try to pass
an [4]int into a function that expects [5]int, it won't compile. They are different types so
it's just the same as trying to pass a string into a function that wants an int.

Go has slices which do not encode the size of the collection and instead can have any size.
if just omit the size during declration it will turn into arrray

Go tool : coverage https://go.dev/blog/cover  , CMD : "go test -cover "
Test coverage is a term that describes how much of a package’s code is exercised by running the
package’s tests. If executing the test suite causes 80% of the package’s source statements to be run,

	we say that the test coverage is 80%.
*/

func TestSum(t *testing.T) {
	t.Run("sum of 5 elements", func(t *testing.T) {
		// using array
		// numbers := [5]int{1,2,3,4,5}
		// using slice
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Got %d wanted %d , given %v", got, want, numbers)
		}
	})
	// redundant test as coverage remains same even when this is uncommented
	// t.Run("sum of as many elements you want", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}
	// 	got := Sum(numbers)
	// 	want := 6
	// 	if got != want {
	// 		t.Errorf("got %d wanted %d , given %v", got, want, numbers)
	// 	}
	// })

}

// Take a varying number of slices, returning a new slice containing the totals for each slice

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	//Go does not let you use equality operators with slices. You could write a function to iterate
	//over each got and want slice and check their values but for convenience sake,
	//we can use reflect.DeepEqual : https://pkg.go.dev/reflect#DeepEqual , also import "reflect"

	// NOTE:  reflect.DeepEqual is not "type safe"  , it will let you compare string and slice too

	// if got != want {
	// 	t.Errorf("got %v want %v", got, want)
	// }
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumsAllTails(t *testing.T) {
	/*

		assigning a function to a variable. , functions in effect are values too.
		this technique is useful when you want to bind a function to other local variables in "scope"
		It also allows you to reduce the surface area of your API.
		By defining this function inside the test, it cannot be used by other functions in this package.
		Hiding variables and functions that don't need to be exported is an important design consideration.
		this adds a little type-safety to our code. If a developer mistakenly adds a new test with
		checkSums(t, got, "dave") the compiler "our friend" will stop them in their tracks.

	*/

	CheckSum := func(t testing.TB, got, want []int) {
		// remember to include t.Helper()
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make sum of tails of sum slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		CheckSum(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		CheckSum(t, got, want)

	})
}

/*

	Compile time errors are our friend because they help us write software that works,
	runtime errors are our enemies because they affect our users.

*/
