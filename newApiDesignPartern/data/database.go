package data

import (
	"encoding/json"
	"net/http"
)

type DB struct {
	Users    []User
	Products []Product
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) ToJsonP(rw http.ResponseWriter) error {
	err := json.NewEncoder(rw).Encode(db.Products)
	return err
}

func (db *DB) AddProduct(p Product) {
	db.Products = append(db.Products, p)
}

func (db *DB) AddUser(u User) {
	db.Users = append(db.Users, u)
}

func (db *DB) GetUserById(id int) (usr User, msg string) {
	if len(db.Users)-1 < id {
		return *NewUser(), "No user found with given Id "
	}
	return db.Users[id], ""
}

func (db *DB) GetProdrById(id int) (p Product, msg string) {
	if len(db.Products)-1 < id {
		return *NewProduct(), "No product found with given Id "
	}
	return db.Products[id], ""
}

var DATABASE = DB{
	Users: []User{
		*NewUserWithFields("loki", 1, "loki@loki", "nill"),
		*NewUserWithFields("sati", 1, "bala@loki", "nill"),
		*NewUserWithFields("bala", 1, "sati@loki", "nill"),
		*NewUserWithFields("muthu", 1, "muthu@loki", "nill"),
	},
	Products: []Product{
		*NewProductWithFields("lays", 34, 10, "404"),
		*NewProductWithFields("bingo", 35, 10, "404"),
		*NewProductWithFields("kutkura", 36, 10, "404"),
		*NewProductWithFields("Pons", 37, 50, "404"),
	},
}
