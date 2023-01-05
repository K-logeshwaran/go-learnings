package handlers

import (
	"log"
	"net/http"
	//"newApiDesignPartern/data"
	//"newApiDesignPartern/data"
)

type UserHandler struct {
	l  *log.Logger
	lf *log.Logger
}

func NewUserHandler(std *log.Logger, lf *log.Logger) *UserHandler {
	return &UserHandler{
		l:  std,
		lf: lf,
	}
}

func (p *UserHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("New Request in /user")

	//rw.Write([]byte("Hi Bro!"))
}
