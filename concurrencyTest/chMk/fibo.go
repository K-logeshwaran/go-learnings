package main

import (
	"fmt"
	//"sync"
)

func fibo(limit int, ch chan int) {
	a := -1
	b := 1

	for i := 0; i < limit; i++ {
		s := a + b
		//fmt.Println(s)
		ch <- s
		a = b
		b = s

	}

	close(ch)

}

func main() {
	//mut := &sync.Mutex{}
	ch := make(chan int, 20)
	go fibo(20, ch)

	for va := range ch {
		fmt.Println(va)
	}
}
