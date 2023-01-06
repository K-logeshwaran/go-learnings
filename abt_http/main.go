package main

import (
	"log"
	"net/url"
)

func main() {
	myUrl := "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

	result, _ := url.Parse(myUrl)

	log.Println(result.Path)
	log.Println(result.Host)
	log.Println(result.Port())
	log.Println(result.RawQuery)
	log.Println(result.RawFragment)
	qPrams := result.Query()

	for key, val := range qPrams {
		log.Println(key, " : ", val)
	}

}
