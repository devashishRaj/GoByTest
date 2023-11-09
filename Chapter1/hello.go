package main

import "fmt"

// It's worth thinking about creating ants to capture the meaning of values and to even aid performance.
const (
	spanish = "Spanish"
	french  = "French"

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

func main() {
	fmt.Println(Hello("World", ""))
}
