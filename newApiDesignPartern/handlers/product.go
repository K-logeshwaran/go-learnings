package handlers

import (
	"io"
	"log"
	"net/http"
	"newApiDesignPartern/data"
	//"newApiDesignPartern/data"
)

type ProductHandler struct {
	l  *log.Logger
	lf *log.Logger
}

func NewProductHandler(std *log.Logger, lf *log.Logger) *ProductHandler {
	return &ProductHandler{
		l:  std,
		lf: lf,
	}
}

func (p *ProductHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		p.getAllProducts(rw)
		p.l.Println("New Get request")

	case http.MethodPost:
		p.l.Println("New Post request")
		p.l.Println(req)
		p.addProducts(req.Body, &data.DATABASE)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(rw, "", http.StatusMethodNotAllowed)
	}

	//rw.Write([]byte("Hi Bro!"))
}

func (ph *ProductHandler) addProducts(json io.Reader, db *data.DB) error {
	p := data.NewProduct()
	err := p.ToStruct(json)
	if err != nil {
		ph.lf.Println("Unable Unmarshal json")
		return err
	}
	db.AddProduct(*p)
	return err
}

func (ph *ProductHandler) getAllProducts(rw http.ResponseWriter) {
	err := data.DATABASE.ToJsonP(rw)
	if err != nil {
		ph.lf.Println("Error")
		http.Error(rw, "cannot Parse json", http.StatusInternalServerError)
	}
}
