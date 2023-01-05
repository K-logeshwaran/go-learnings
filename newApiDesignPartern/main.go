package main

import (
	//"fmt"
	"log"
	"net"
	"net/http"
	"os"

	//"net/http"
	//"newApiDesignPartern/data"
	"newApiDesignPartern/handlers"
)

func main() {
	// creating loggers

	l := log.New(os.Stdout, "product Api", log.LstdFlags)
	f, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	lf := log.New(f, "product Api ", log.LstdFlags)

	// New server Mux
	mux := http.NewServeMux()

	// creating  handler Obects  for the paths
	ph := handlers.NewProductHandler(l, lf)
	uh := handlers.NewUserHandler(l, lf)

	// handlers
	go func() {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			http.ServeFile(w, r, "./index.html")
		})
		mux.Handle("/products", ph)
		mux.Handle("/user", uh)
	}()
	ln, err := net.Listen("tcp", "192.168.1.34:8080")
	err2 := http.Serve(ln, mux)

	if err != nil {
		l.Fatal(err2)
	}

	if err2 != nil {
		l.Fatal(err2)
	}

}
