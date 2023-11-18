package countdown

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	// testing what is being printed
	t.Run("sleep and write operations , sequece unchecked", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &SpyCoutdownOperations{})

		got := buffer.String()
		want :=
			`3
2
1
Go!`

		if got != want {
			t.Errorf("Got %q wanted %q ", got, want)
		}
	})
	t.Run("sequece of sleep and write is checked ", func(t *testing.T) {
		spySleepPrinter := &SpyCoutdownOperations{}
		Countdown(spySleepPrinter ,spySleepPrinter )

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want ,spySleepPrinter.Calls){
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}

	})

}

func TestConfigurableSleep(t *testing.T){
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime ,spyTime.Sleep}
	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime{
		t.Errorf("should have slept for %v but slept for %v " , sleepTime , spyTime.durationSlept)
	}
}

/*
If your mocking code is becoming complicated or you are having to mock out lots of things to test something, 
you should listen to that bad feeling and think about your code. Usually it is a sign of
The thing you are testing is having to do too many things (because it has too many dependencies to mock)
Break the module apart so it does less
Its dependencies are too fine-grained
Think about how you can consolidate some of these dependencies into one meaningful module
Your test is too concerned with implementation details
Favour testing expected behaviour rather than the implementation
Normally a lot of mocking points to bad abstraction in your code.

Ever run into this situation?
You want to do some refactoring
To do this you end up changing lots of tests
You question TDD and make a post on Medium titled "Mocking considered harmful"

This is usually a sign of you testing too much implementation detail. Try to make it so your tests 
are testing useful behaviour unless the implementation is really important to how the system runs.
It is sometimes hard to know what level to test exactly but here are some thought processes and 
rules I try to follow:
The definition of refactoring is that the code changes but the behaviour stays the same. 
If you have decided to do some refactoring in theory you should be able to make the commit without 
any test changes. So when writing a test ask yourself
Am I testing the behaviour I want, or the implementation details?
If I were to refactor this code, would I have to make lots of changes to the tests?
Although Go lets you test private functions, I would avoid it as private functions are implementation 
detail to support public behaviour. Test the public behaviour. Sandi Metz describes private functions
 as being "less stable" and you don't want to couple your tests to them.
I feel like if a test is working with more than 3 mocks then it is a red flag - time for a rethink on the design
Use spies with caution. Spies let you see the insides of the algorithm you are writing which can
 be very useful but that means a tighter coupling between your test code and the implementation. 
 Be sure you actually care about these details if you're going to spy on them

*/
