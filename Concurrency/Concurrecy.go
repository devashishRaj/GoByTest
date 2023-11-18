// http://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast

package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
//This loop launches a goroutine for each URL in the urls slice. Each goroutine executes the wc 
//function (assumed to be a function that checks the status of a website) and 
//sends the result to the resultChannel.
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}

		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

/*
Sometimes, when we run our tests, two of the goroutines write to the results map at exactly the same time.
Maps in Go don't like it when more than one thing tries to write to them at once, and so fatal error.
This is a race condition, a bug that occurs when the output of our software is dependent on the
timing and sequence of events that we have no control over. Because we cannot control exactly when
each goroutine writes to the results map, we are vulnerable to two goroutines writing to it at the
same time.


-goroutines, the basic unit of concurrency in Go, which let us check more than one website at the same time.
-anonymous functions, which we used to start each of the concurrent processes that check websites.
-channels, to help organize and control the communication between the different processes, allowing 
us to avoid a race condition bug.
-the race detector which helped us debug problems with concurrent code

*/
