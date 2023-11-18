// NOTE : https://martinfowler.com/bliki/TestDouble.html
package countdown

import (
	"fmt"
	"io"
	"time"
)

const (
	finalword = "Go!"
	countdown = 3
	write = "write"
	sleep = "sleep"
)

type Sleeper interface{
	Sleep()
}

type SpySleeper struct{
	Calls int
}

type DefaultSleeper struct{}
func (d *DefaultSleeper) Sleep(){
	time.Sleep(1*time.Second)
}

type SpyCoutdownOperations struct{
	Calls []string
}

func (s *SpyCoutdownOperations)Sleep(){
	s.Calls = append(s.Calls, sleep )
}

func (s *SpyCoutdownOperations) Write(p []byte)(n int , err error){
	s.Calls = append(s.Calls, write)
	return
}

/*
	Spies are a kind of mock which can record how a dependency is used. They can record the 
	arguments sent in, how many times it has been called, etc. In our case, we're keeping track of 
	how many times Sleep() is called so we can check it in our test.
*/
func (s *SpySleeper) Sleep(){
	s.Calls++
}


type ConfigurableSleeper struct{
	duration 	time.Duration
	sleep 		func(time.Duration)
}

type SpyTime struct{
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration){
	s.durationSlept = duration
}
func(c *ConfigurableSleeper) Sleep(){
	c.sleep(c.duration)
}
// func Countdown(out io.Writer) {
// 	// We're using fmt.Fprint which takes an io.Writer (like *bytes.Buffer) and sends a string to it.
// 	for i := countdown; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 		// Our tests take 3 seconds to run , BAD!
// 		// Slow tests ruin developer productivity.
// 		time.Sleep(1 * time.Second)
// 	}
// 	fmt.Fprint(out, finalword)
// }

/*
	If we can mock time.Sleep we can use dependency injection to use it instead of a 
	"real" time.Sleep and then we can spy on the calls to make assertions on them.


*/


// Our latest change only asserts that it has slept 3 times, but those sleeps could occur out of sequence.
func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdown; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalword)
}



