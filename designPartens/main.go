package main

import (
	"log"
	"net/http"
	"os"

	"github.com/designPartens/handlers"
)

func main() {
	file, err := os.OpenFile("custom.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	l := log.New(file, "products-api ", log.LstdFlags)
	sm := http.NewServeMux()
	hh := handlers.NewHello(l)
	sm.HandleFunc("/", hh.ServeHTTP)
	http.ListenAndServe(":5000", sm)
}
