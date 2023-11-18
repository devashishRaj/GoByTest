package selectRacer

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var timeout = 10 * time.Millisecond

// net/http/httptest enables users to easily create a mock HTTP server.
// you'll want to include time.After in one of your cases to prevent your system blocking forever.

func Racer(a, b string) (winner string, err error) {
	/*
		you can wait for values to be sent to a channel with myVar := <-ch. This is a blocking call,
		as you're waiting for a value.
		select allows you to wait on multiple channels. The first one to send a value "wins" and the code
		underneath the case is executed.
		We use ping in our select to set up two channels, one for each of our URLs. Whichever one writes
		to its channel first will have its code executed in the select, which results in its URL being
		returned (and being the winner).

	*/
	return configurableRacer(a, b, timeout)

}

func configurableRacer(a, b string, servertimeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(servertimeout):
		return "", fmt.Errorf("timed out wating for %s ans %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// always use make for channels as var will lead to nil/zero value
	//For channels the zero value is nil and if you try and send to it with <- it will block forever
	// because you cannot send to nil channels
	ch := make(chan struct{})
	go func() {
		_, err := http.Get(url)
		log.Fatalln(err)
		close(ch)
	}()
	return ch
}

// func measureTime(url string) time.Duration {
// 	start := time.Now()
// 	_ , err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return time.Since(start)
// }
