package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}
	got := CheckWebsites(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v wanted %v", got, want)
	}

}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//CheckWebsites using a slice of one hundred urls and uses a new fake implementation of
		//WebsiteChecker. slowStubWebsiteChecker is deliberately slow. It uses time.Sleep to wait
		//exactly twenty milliseconds and then it returns true. We use b.ResetTimer() in this test to
		//reset the time of our test before it actually runs
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

/*

	Anonymous functions have a number of features which make them useful, two of which we're using above.
	Firstly, they can be executed at the same time that they're declared - this is what the () at
	the end of the anonymous function is doing. Secondly they maintain access to the lexical scope
	in which they are defined - all the variables that are available at the point when you declare
	the anonymous function are also available in the body of the function.
*/
