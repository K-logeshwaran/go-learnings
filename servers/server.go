package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Books struct {
	Name      string     `json:"name"`
	Ratings   float64    `json:"ratings"`
	Author    string     `json:"author"`
	Price     float64    `json:"price"`
	Publisher Publishers `json:"publisher"`
}

type Publishers struct {
	Name        string `json:"name"`
	Contact     string `json:"contact"`
	No_of_books int    `json:"no_of_books"`
}

var Database = []Books{
	{Name: "Automate", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
	{Name: "Golang", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
	{Name: "Node js", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
	{Name: "Rust", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
	{Name: "Bot", Ratings: 9.9, Author: "Don't kn", Price: 2000.75, Publisher: Publishers{Name: "Idk", Contact: "dkdmdmakdmdd", No_of_books: 34}},
}

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		// Serve the resource.
		id := req.FormValue("id")
		println(id)
		if id == "" {
			D, _ := json.Marshal(Database)
			w.Write(D)
			return
		}
		intId, _ := strconv.Atoi(id)
		if intId > len(Database) {
			http.Error(w, "No records Found at id "+string(id), http.StatusNotFound)
			return
		}
		D, _ := json.Marshal(Database[intId])
		w.Write(D)
	case http.MethodPost:
		var newRecord Books
		err := json.NewDecoder(req.Body).Decode(&newRecord)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		Database = append(Database, newRecord)
		w.Write([]byte("Data added Successfully"))
		fmt.Println(Database)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// // Set the Content-Type header so the browser knows what kind of file it is
	// w.Header().Set("Content-Type", "application/octet-stream")
	// // Set the Content-Disposition header so the browser presents the option to save the file
	// //w.Header().Set("Content-Disposition", "attachment; filename=myfile.txt")
	// // Serve the file using the ServeFile function
	// http.ServeFile(w, r, "./index.txt")
	f, err := os.Open("./index.txt")
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/octet-stream")
	// Set the Content-Disposition header
	w.Header().Set("Content-Disposition", "attachment; filename=myfile.txt")

	// Copy the contents of the file to the response writer
	_, err = io.Copy(w, f)
	if err != nil {
		http.Error(w, "Error copying file", http.StatusInternalServerError)
		return
	}
}

func main() {
	fmt.Println("Hello vim you are my friend now \n It's  time to make a change")

	fmt.Println(Database[:])
	//rm:=append(Database,Database[:4])
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/download", downloadHandler)

	err := http.ListenAndServe(":8081", mux)

	log.Fatal(err)
}
