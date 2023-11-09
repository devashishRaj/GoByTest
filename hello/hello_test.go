package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello world to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello , Chris"

		assertCorrectMsg(t, got, want)
		//group tests around a "thing" and then have subtests describing different scenarios.
	})
	t.Run("say 'Hello , World' when an empty string is passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello , World"

		assertCorrectMsg(t, got, want)
	})

}

// reduces duplication and improves readability of our tests.
//	pass in t *testing.T so that we can tell the test code to fail when we need to.
// accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy,
//so you can call helper functions from a test, or a benchmark
/*
t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it
fails the line number reported will be in our function call rather than inside our test helper.
This will help other developers track down problems easier.
it will tell the line where test failed other wise it will report line number of this assert func
*/
func assertCorrectMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
