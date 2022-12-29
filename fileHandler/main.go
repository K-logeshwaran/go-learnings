package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	contend := "My first file handler "
	file, err := os.Create("./hello.txt")
	defer file.Close()
	checkPanic(err)
	result, err := io.WriteString(file, contend)
	checkPanic(err)
	fmt.Println(result)
}

func checkPanic(e error) {
	if e != nil {
		panic(e)
	}
}
