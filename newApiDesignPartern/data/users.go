package data

import (
	"encoding/json"

	"net/http"
	//"net/http"
)

type User struct {
	Name     string `json:"name"`
	Id       int    `json:"productId"`
	Email    string `json:"emailId"`
	Password string `json:"-"`
}

func (p *User) ToJson(w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(p)
	return err
}

func (p *User) ToStruct(rq *http.Request) error {
	err := json.NewDecoder(rq.Body).Decode(p)
	return err
}

func NewUser() *User {
	return &User{}
}

func NewUserWithFields(name string, id int, email string, password string) *User {
	return &User{
		Name:     name,
		Id:       id,
		Email:    email,
		Password: password,
	}
}
