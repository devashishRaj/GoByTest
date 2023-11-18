package dependency

import (
	"fmt"
	"io"
)

// bytes.Buffer  is not general purpose interface, thus you can use it print to stdout like
// in a go application like this : Greet(os.Stdout, "Elodie")
// it will lead to error
// cannot use os.Stdout (type *os.File) as type *bytes.Buffer in argument to Greet

// func Greet(writer *bytes.Buffer, name string) {
// 	fmt.Fprintf(writer,"Hello, %s", name)
// }

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
