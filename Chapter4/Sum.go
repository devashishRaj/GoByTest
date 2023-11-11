package arraySlices

func Sum(numbers []int) int {
	sum := 0
	// range lets you iterate over an array. On each iteration, range returns two values - the index
	// and the value. We are choosing to ignore the index value by using _
	//https://go.dev/doc/effective_go#blank
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*
slices have a capacity. If you have a slice with a capacity of 2 and try to do mySlice[10] = 1
you will get a runtime error.
*/
func SumAll(numbersToSum ...[]int) []int {
	// length := len(numbersToSum)
	// // make allows you to create a slice with a starting capacity of the len of the numbersToSum
	// sums := make([]int, length)
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	var sums []int
	for _, numbers := range numbersToSum {
		//you can use the append function which takes a slice and a new value,
		//then returns a new slice with all the items in it.
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		//Slices can be sliced! The syntax is slice[low:high]. If you omit the value on one of the
		//sides of the : it captures everything to that side of it. In our case,
		// we are saying "take from 1 to the end" with numbers[1:].
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
