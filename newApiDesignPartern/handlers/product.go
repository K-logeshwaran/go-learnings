package handlers

import (
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

func (ph *ProductHandler) AddProducts(rw http.ResponseWriter, req *http.Request) {
	p := data.NewProduct()
	err := p.ToStruct(req.Body)
	if err != nil {
		ph.lf.Println("Unable Unmarshal json")
		http.Error(rw, "Unable Unmarshal json", http.StatusInternalServerError)
		return
	}
	data.DATABASE.AddProduct(*p)
}

func (ph *ProductHandler) GetAllProducts(rw http.ResponseWriter, req *http.Request) {
	err := data.DATABASE.ToJsonP(rw)
	if err != nil {
		ph.lf.Println("Error")
		http.Error(rw, "cannot Parse json", http.StatusInternalServerError)
	}
}
