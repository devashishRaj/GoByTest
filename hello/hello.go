package main

import "fmt"

//It's worth thinking about creating constants to capture the meaning of values and to even aid performance.
const greetingPrefix = "Hello , "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix + name
}

func main(){
	fmt.Println(Hello("World"))
}
