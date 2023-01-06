package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	// Read up to 5 bytes from the reader into a slice.
	value, err := reader.ReadString('.')
	if err != nil {
		// Handle the error.
		fmt.Println(err)
		return
	}

	// Print the slice that was read.
	fmt.Printf("%s\n", value)
	fmt.Printf("%v\n",'.')
	println(value)
}
