package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"
	if repeated != expected {
		t.Errorf("Expected %q got %q", expected, repeated)
	}
}

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		// ADD and Asginment operatro "+="
		repeated += character
	}
	return repeated
}

/*

other varaints of for loop


func forvariants() {

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}

*/

// writing benchmark
// To run the benchmarks do " go test -bench=. "
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

//What 136 ns/op means is our function takes on average 136 nanoseconds to run (on my computer).

// NOTE by default Benchmarks are run sequentially.

func ExampleRepeat() {
	repeatSix := Repeat("b", 6)
	fmt.Println(repeatSix)
	// Output: bbbbbb
}
