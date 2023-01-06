package main

import "fmt"

func main() {
	a := -1
	b := 1

	for i := 0; i < 20; i++ {
		s := a + b
		fmt.Println(s)
		a = b
		b = s

	}

}
