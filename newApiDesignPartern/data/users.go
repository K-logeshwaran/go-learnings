package data

import (
	"encoding/json"
	"io"

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

func (p *User) ToStruct(data io.Reader) error {
	err := json.NewDecoder(data).Decode(p)
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
