package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var Wg sync.WaitGroup

func main() {
	var sites = []string{
		"https://instagram.com/",
		"https://google.com/",
		"https://github.com/",
	}

	for _, site := range sites {
		go getStatus(site)
		//getStatus(site)
		Wg.Add(1)
	}
	Wg.Wait()
}

func getStatus(site string) {
	defer Wg.Done()
	res, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response status %s  from %s \n", res.Status, site)
}
