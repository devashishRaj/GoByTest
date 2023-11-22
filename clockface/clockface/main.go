package main

import (
	"os"
	"time"

	clockface "gobyTests/clockface" // REPLACE THIS!
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
