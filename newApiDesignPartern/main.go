package main

import (
	//"fmt"
	"log"
	"net"
	"net/http"
	"os"

	//"net/http"
	//"newApiDesignPartern/data"
	"newApiDesignPartern/data"
	"newApiDesignPartern/handlers"

	"github.com/gorilla/mux"
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
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	// creating  handler Obects  for the paths
	ph := handlers.NewProductHandler(l, lf)
	uh := handlers.NewUserHandler(l, lf)

	// handlers
	go func() {

		sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			http.ServeFile(w, r, "./index.html")
		})
		getRouter.HandleFunc("/products", ph.GetAllProducts)
		postRouter.HandleFunc("/products", ph.AddProducts)

		getRouter.HandleFunc("/user", uh.GetAllUsers)
		getRouter.HandleFunc("/user/{id}", uh.GetUserById)
		postRouter.HandleFunc("/user", uh.AddUsers)
		l.Println(data.DATABASE)
	}()
	ln, err := net.Listen("tcp", "192.168.1.34:8080")
	err2 := http.Serve(ln, sm)

	if err != nil {
		l.Fatal(err2)
	}

	if err2 != nil {
		l.Fatal(err2)
	}

}
