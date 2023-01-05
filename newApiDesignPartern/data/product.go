package data

import (
	"encoding/json"
	"io"
	"net/http"
	//"net/http"
)

type Product struct {
	Name  string  `json:"name"`
	Id    int     `json:"productId"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

func (p *Product) ToJson(w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(p)
	return err
}

func (p *Product) ToStruct(r io.Reader) error {
	err := json.NewDecoder(r).Decode(p)
	return err
}

func NewProduct() *Product {
	return &Product{}
}

func NewProductWithFields(name string, id int, price float64, image string) *Product {
	return &Product{
		Name:  name,
		Id:    id,
		Price: price,
		Image: image,
	}
}
