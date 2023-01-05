package handlers

import (
	"log"
	"net/http"
	"newApiDesignPartern/data"
	"strconv"

	"github.com/gorilla/mux"
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

func (ph *UserHandler) AddUsers(rw http.ResponseWriter, req *http.Request) {
	p := data.NewUser()
	err := p.ToStruct(req.Body)
	if err != nil {
		ph.lf.Println("Unable Unmarshal json")
		http.Error(rw, "Unable Unmarshal json", http.StatusInternalServerError)
		return
	}
	data.DATABASE.AddUser(*p)
}

func (ph *UserHandler) GetAllUsers(rw http.ResponseWriter, req *http.Request) {
	err := data.DATABASE.ToJsonU(rw)
	if err != nil {
		ph.lf.Println("Error")
		http.Error(rw, "cannot Parse json", http.StatusInternalServerError)
	}
}

func (ph *UserHandler) GetUserById(rw http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)
	ph.l.Println(id)
	intId, _ := strconv.Atoi(id["id"])
	usr, msg := data.DATABASE.GetUserById(intId)
	if msg != "" {
		ph.lf.Println("Error")
		rw.Write([]byte(msg))
		return
	}
	err := usr.ToJson(rw)
	if err != nil {
		ph.lf.Println("Error")
		http.Error(rw, "cannot Parse json", http.StatusInternalServerError)
	}
}
