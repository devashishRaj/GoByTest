package main

import (
	countdown "gobyTests/Countdown"
	
	"os"
	"time"
)

// It's worth thinking about creating ants to capture the meaning of values and to even aid performance.
const (
	//spanish = "Spanish"
	//french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// (prefix string) This will create a variable called prefix in your function. default 0 or ""
// You can return whatever it's set to by just calling return rather than return prefix.
// In Go, public functions start with a capital letter and private ones start with a lowercase.
func greetingPrefix(language string) (prefix string) {

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper is an implementation of Sleeper with a defined delay.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep will pause execution for the defined Duration.
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	//fmt.Println(Hello("World", ""))
	//dependency.Greet(os.Stdout, "Elodie")

	// chapter 8 : mocking
	//sleeper := countdown.ConfigurableSleeper{1 * time.Second, time.Sleep}
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	countdown.Countdown(os.Stdout, sleeper)


}
